package team

import (
	"context"
	"errors"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/cache"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/repository"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/google/uuid"
)

//go:generate mockgen -destination ../../../mock/team_mock.go -package=mock -source=service.go
type Team interface {
	RegisterTeam(ctx context.Context, request *dto.TeamRequest) (*dto.TeamResponse, error)
	UpdateTeam(ctx context.Context, id string, request *dto.TeamRequest) error
	GetTeamByID(ctx context.Context, id string) (*dto.TeamResponse, error)
	GetTeamByName(ctx context.Context, name *dto.GetTeamByName) (*dto.TeamResponse, error)
	DeleteTeamByID(ctx context.Context, id string) error
}

type service struct {
	repository repository.Repository
	cache      cache.Cache
	log        log.Log
}

func New(repository repository.Repository, cache cache.Cache,
	log log.Log) *service {
	return &service{
		repository: repository,
		cache:      cache,
		log:        log,
	}
}

func (s *service) RegisterTeam(ctx context.Context, request *dto.TeamRequest) (*dto.TeamResponse, error) {
	if !model.CheckUniverse(model.Universe(request.Universe)) {
		return nil, exception.ErrInvalidFields
	}
	team := model.NewTeam(uuid.NewString(),
		time.Now().UTC(), request)
	if err := s.repository.CreateTeam(ctx, team); err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}

	if err := s.cache.SetTeam(ctx, team, team.ID); err != nil {
		s.log.SendErrorLog(ctx, err)
	}

	return dto.NewTeamResponse(team.ID, team.Name, team.Universe,
		team.CreatedAt, team.UpdatedAt), nil
}

func (s *service) GetTeamByID(ctx context.Context, id string) (*dto.TeamResponse, error) {
	team, err := s.cache.GetTeam(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		if team, err = s.repository.GetTeamByID(ctx, id); err != nil {
			s.log.SendErrorLog(ctx, err)
			return nil, exception.ErrTeamNotFound
		}

	}
	return dto.NewTeamResponse(
		team.ID,
		team.Name,
		team.Universe,
		team.CreatedAt,
		team.UpdatedAt), nil
}

func (s *service) DeleteTeamByID(ctx context.Context, id string) (err error) {
	if err = s.cache.DeleteTeam(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err)
		return exception.ErrInternalServer
	}
	if err = s.repository.DeleteTeamByID(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err)
		if errors.Is(err, exception.ErrTeamNotFound) {
			return
		}
		return exception.ErrInternalServer
	}

	return nil
}

func (s *service) GetTeamByName(ctx context.Context,
	request *dto.GetTeamByName) (*dto.TeamResponse, error) {
	team, err := s.cache.GetTeam(ctx, request.Name)
	if err != nil {
		team, err = s.repository.GetTeamByName(ctx, request.Name)
		if err != nil {
			s.log.SendErrorLog(ctx, err)
			return nil, exception.ErrTeamNotFound
		}

	}
	if err := s.cache.SetTeam(ctx, team, request.Name); err != nil {
		s.log.SendErrorLog(ctx, err)
	}
	return dto.NewTeamResponse(
		team.ID,
		team.Name,
		team.Universe,
		team.CreatedAt,
		team.UpdatedAt,
	), nil
}

func (s *service) UpdateTeam(ctx context.Context, id string,
	dto *dto.TeamRequest) error {
	team := model.NewTeam(id, time.Now(), dto)
	if !model.CheckUniverse(model.Universe(team.Universe)) {
		s.log.SendErrorLog(ctx, errors.New("invalid field"))
		return exception.ErrInvalidFields
	}
	team.UpdatedAt = util.GerPointer(time.Now().UTC())
	if err := s.repository.UpdateTeam(ctx, team); err != nil {
		s.log.SendErrorLog(ctx, err)
		if err := s.cache.DeleteHero(ctx, team.ID); err != nil {
			s.log.SendErrorLog(ctx, err)
			return exception.ErrInternalServer
		}
		return err
	}

	if err := s.cache.SetTeam(ctx, team, team.ID); err != nil {
		s.log.SendErrorLog(ctx, err)
		if err := s.cache.DeleteHero(ctx, team.ID); err != nil {
			s.log.SendErrorLog(ctx, err)
		}
	}

	if err := s.cache.SetTeam(ctx, team, team.Name); err != nil {
		s.log.SendErrorLog(ctx, err)
		if err := s.cache.DeleteHero(ctx, team.Name); err != nil {
			s.log.SendErrorLog(ctx, err)
		}
	}

	return nil
}
