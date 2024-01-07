package team

import (
	"context"
	"fmt"
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
	CreateTeam(ctx context.Context, request *dto.TeamRequest) (*dto.TeamResponse, error)
	UpdateTeam(ctx context.Context, id string, request *dto.TeamRequest) error
	GetTeamByID(ctx context.Context, id string) (*dto.TeamResponse, error)
	GetTeamByName(ctx context.Context, name *dto.GetTeamByName) (*dto.TeamResponse, error)
	DeleteTeamByID(ctx context.Context, id string) error
}

type service struct {
	repository repository.TeamRepository
	cache      cache.TeamCache
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
func (s *service) CreateTeam(ctx context.Context, req *dto.TeamRequest) (*dto.TeamResponse, error) {
	resp, err := s.createTeam(ctx, req)
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}
	return resp, nil
}

func (s *service) createTeam(ctx context.Context, request *dto.TeamRequest) (*dto.TeamResponse, error) {
	if !model.CheckUniverse(model.Universe(request.Universe)) {
		return nil, exception.New(fmt.Sprintf("checkUniverse\n%s", exception.ErrInvalidFields), exception.ErrInvalidFields)
	}
	team := model.NewTeam(uuid.NewString(), time.Now().UTC(), request)
	if err := s.repository.CreateTeam(ctx, team); err != nil {
		return nil, exception.New(fmt.Sprintf("createTeam\n%s", err.Error()), err)
	}

	if err := s.cache.SetTeam(ctx, team, team.ID); err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("setTeam\n%w", err))
	}

	return dto.NewTeamResponse(team.ID, team.Name, team.Universe,
		team.CreatedAt, team.UpdatedAt), nil
}

func (s *service) GetTeamByID(ctx context.Context, id string) (*dto.TeamResponse, error) {
	resp, err := s.getTeamByID(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}
	return resp, nil
}

func (s *service) getTeamByID(ctx context.Context, id string) (*dto.TeamResponse, error) {
	team, err := s.cache.GetTeam(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("getTeam\n%w", err))
		if team, err = s.repository.GetTeamByID(ctx, id); err != nil {
			return nil, exception.New(fmt.Sprintf("getTeamByID\n%s", err.Error()), exception.ErrTeamNotFound)
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
	if err := s.deleteTeamByID(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err)
		return err
	}
	return nil
}

func (s *service) deleteTeamByID(ctx context.Context, id string) (err error) {
	if err = s.cache.DeleteTeam(ctx, id); err != nil {
		return exception.New(fmt.Sprintf("deleteTeam\n%s", err.Error()), err)
	}

	if err = s.repository.DeleteTeamByID(ctx, id); err != nil {
		return exception.New(fmt.Sprintf("deleteTeamByID\n%s", err.Error()), err)
	}

	return nil
}
func (s *service) GetTeamByName(ctx context.Context,
	req *dto.GetTeamByName) (*dto.TeamResponse, error) {
	resp, err := s.getTeamByName(ctx, req)
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}
	return resp, nil
}

func (s *service) getTeamByName(ctx context.Context,
	req *dto.GetTeamByName) (*dto.TeamResponse, error) {
	team, err := s.cache.GetTeam(ctx, req.Name)
	if err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("getTeam\n%w", err))
		team, err = s.repository.GetTeamByName(ctx, req.Name)
		if err != nil {
			return nil, exception.New(fmt.Sprintf("getTeamByName\n%s", err.Error()), exception.ErrTeamNotFound)
		}

	}
	if err := s.cache.SetTeam(ctx, team, req.Name); err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("setTeam\n%w", err))
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
	req *dto.TeamRequest) error {
	if err := s.updateTeam(ctx, id, req); err != nil {
		s.log.SendErrorLog(ctx, err)
		return err
	}
	return nil
}
func (s *service) updateTeam(ctx context.Context, id string,
	req *dto.TeamRequest) error {
	team := model.NewTeam(id, time.Now(), req)
	if !model.CheckUniverse(model.Universe(team.Universe)) {
		return exception.ErrInvalidFields
	}
	team.UpdatedAt = util.GerPointer(time.Now().UTC())
	if err := s.repository.UpdateTeam(ctx, team); err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("updateTeam\n%w", err))
		if err := s.cache.DeleteTeam(ctx, team.ID); err != nil {
			return exception.New(fmt.Sprintf("deleteTeam\n%s", err.Error()), err)
		}
		return err
	}

	if err := s.cache.SetTeam(ctx, team, team.ID); err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("setTeam\n%w", err))
		if err := s.cache.DeleteTeam(ctx, team.ID); err != nil {
			s.log.SendErrorLog(ctx, fmt.Errorf("deleteTeam\n%w", err))
		}
	}

	if err := s.cache.SetTeam(ctx, team, team.Name); err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("setTeam\n%w", err))
		if err := s.cache.DeleteTeam(ctx, team.Name); err != nil {
			s.log.SendErrorLog(ctx, fmt.Errorf("deleteTeam\n%w", err))
		}
	}

	return nil
}
