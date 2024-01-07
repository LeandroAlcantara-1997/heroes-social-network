package game

import (
	"context"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/cache"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/repository"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/google/uuid"
)

type Game interface {
	Create(ctx context.Context, req *dto.GameRequest) (*dto.GameResponse, error)
	UpdateGame(ctx context.Context, id string, req *dto.GameRequest) error
	GetByID(ctx context.Context, id string) (*dto.GameResponse, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository repository.GameRepository
	cache      cache.GameCache
	log        log.Log
}

func New(repository repository.GameRepository, cache cache.GameCache, log log.Log) *service {
	return &service{
		repository: repository,
		cache:      cache,
		log:        log,
	}
}

func (s *service) Create(ctx context.Context, req *dto.GameRequest) (*dto.GameResponse, error) {
	var game = model.NewGame(uuid.NewString(), req)
	if err := s.repository.CreateGame(ctx, game); err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}

	if err := s.cache.SetGame(ctx, game); err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}

	return &dto.GameResponse{
		ID:          game.ID,
		Name:        game.Name,
		ReleaseYear: game.ReleaseYear,
		Universe:    game.Universe,
	}, nil
}

func (s *service) UpdateGame(ctx context.Context, id string, req *dto.GameRequest) error {
	game := model.NewGame(id, req)
	game.UpdatedAt = util.GerPointer(time.Now().UTC())
	if err := s.repository.UpdateGame(ctx, game); err != nil {
		s.log.SendErrorLog(ctx, err)
		return err
	}

	if err := s.cache.SetGame(ctx, game); err != nil {
		s.log.SendErrorLog(ctx, err)
		if err := s.cache.DeleteGame(ctx, game.ID); err != nil {
			s.log.SendErrorLog(ctx, err)
		}
	}

	return nil
}

func (s *service) GetByID(ctx context.Context, id string) (*dto.GameResponse, error) {
	game, err := s.cache.GetGame(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		game, err = s.repository.GetGameByID(ctx, id)
		if err != nil {
			s.log.SendErrorLog(ctx, err)
			return nil, err
		}
	}
	return &dto.GameResponse{
		ID:          game.ID,
		Name:        game.Name,
		ReleaseYear: game.ReleaseYear,
		Universe:    game.Universe,
		TeamID:      game.TeamID,
		HeroID:      game.HeroID,
		CreatedAt:   &game.CreatedAt,
		UpdatedAt:   game.UpdatedAt,
	}, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	if err := s.cache.DeleteGame(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err)
		return err
	}

	if err := s.repository.DeleteGameByID(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err)
		return err
	}
	return nil
}
