package game

import (
	"context"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/cache"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/repository"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/model"
	"github.com/google/uuid"
)

type Game interface {
	Create(ctx context.Context, req *dto.GameRequest) (*dto.GameResponse, error)
}

type service struct {
	repository repository.GameRepository
	cache      cache.GameCache
}

func New(repository repository.GameRepository, cache cache.GameCache) *service {
	return &service{
		repository: repository,
		cache:      cache,
	}
}

func (s *service) Create(ctx context.Context, req *dto.GameRequest) (*dto.GameResponse, error) {
	var game = model.NewGame(uuid.NewString(), req)
	if err := s.repository.CreateGame(ctx, game); err != nil {
		return nil, err
	}

	if err := s.cache.SetGame(ctx, game, game.ID); err != nil {
		return nil, err
	}

	return &dto.GameResponse{
		ID:          game.ID,
		Name:        game.Name,
		ReleaseYear: game.ReleaseYear,
		Universe:    game.Universe,
	}, nil
}

func (s *service) Update(ctx context.Context, req *dto.GameRequest) (*dto.GameResponse, error) {
	return nil, nil
}

func (s *service) GetByID(ctx context.Context, req *dto.GameRequest) (*dto.GameResponse, error) {
	return nil, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	return nil
}
