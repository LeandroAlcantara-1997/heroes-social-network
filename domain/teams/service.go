package teams

import (
	"context"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
	input "github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/output/cache"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/output/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/output/repository"
	"github.com/google/uuid"
)

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

func (s *service) RegisterTeam(ctx context.Context, request *input.TeamRequest) (*input.TeamResponse, error) {
	if !model.CheckUniverse(model.Universe(request.Universe)) {
		return nil, exception.ErrInvalidFields
	}
	team := model.NewTeam(uuid.NewString(), time.Now().UTC(), request)
	if err := s.repository.CreateTeam(ctx, team); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		return nil, exception.ErrInternalServer
	}

	if err := s.cache.SetTeam(ctx, team); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
	}

	return input.NewTeamResponse(team.ID, team.Name, team.Universe,
		team.CreatedAt, team.UpdatedAt), nil
}

func (s *service) GetTeamByID(ctx context.Context, id string) (*input.TeamResponse, error) {
	team, err := s.cache.GetTeam(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		if team, err = s.repository.GetTeamByID(ctx, id); err != nil {
			s.log.SendErrorLog(ctx, err.Error())
			return nil, exception.ErrTeamNotFound
		}

	}
	return input.NewTeamResponse(
		team.ID,
		team.Name,
		team.Universe,
		team.CreatedAt,
		team.UpdatedAt), nil
}
