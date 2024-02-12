package game

import (
	"context"
	"fmt"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/cache"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/repository"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/google/uuid"
)

//go:generate mockgen -destination ../../../mock/game_mock.go -package=mock -source=service.go
type Game interface {
	CreateGame(ctx context.Context, req *dto.GameRequest) (*dto.GameResponse, error)
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
func (s *service) CreateGame(ctx context.Context, req *dto.GameRequest) (*dto.GameResponse, error) {
	resp, err := s.createGame(ctx, req)
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) createGame(ctx context.Context, req *dto.GameRequest) (*dto.GameResponse, error) {
	var game = model.NewGame(uuid.NewString(), req)
	if err := s.repository.CreateGame(ctx, game); err != nil {
		return nil, exception.New(fmt.Sprintf("createGame\n%s", err.Error()), err)
	}

	if err := s.cache.SetGame(ctx, game); err != nil {
		return nil, exception.New(fmt.Sprintf("setGame\n%s", err.Error()), err)
	}

	return dto.NewGameResponse(game.ID, game.Name, game.ReleaseYear,
		game.TeamID, game.Universe, game.CreatedAt,
		game.UpdatedAt, game.HeroID, game.Consoles), nil
}

func (s *service) UpdateGame(ctx context.Context, id string, req *dto.GameRequest) error {
	if err := s.updateGame(ctx, id, req); err != nil {
		s.log.SendErrorLog(ctx, err)
		return err
	}

	return nil
}

func (s *service) updateGame(ctx context.Context, id string, req *dto.GameRequest) error {
	game := model.NewGame(id, req)
	game.UpdatedAt = util.GerPointer(time.Now().UTC())
	if err := s.repository.UpdateGame(ctx, game); err != nil {
		return exception.New(fmt.Sprintf("updateGame\n%s", err.Error()), err)
	}

	if err := s.cache.SetGame(ctx, game); err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("setGame\n%w", err))
		if err := s.cache.DeleteGame(ctx, game.ID); err != nil {
			s.log.SendErrorLog(ctx, fmt.Errorf("deleteGame\n%w", err))
		}
	}

	return nil
}
func (s *service) GetByID(ctx context.Context, id string) (*dto.GameResponse, error) {
	resp, err := s.getByID(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}
	return resp, nil
}

func (s *service) getByID(ctx context.Context, id string) (*dto.GameResponse, error) {
	game, err := s.cache.GetGame(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("getGame\n%w", err))
		game, err = s.repository.GetGameByID(ctx, id)
		if err != nil {
			return nil, exception.New(fmt.Sprintf("getGameByID\n%s", err.Error()), err)
		}
	}
	return dto.NewGameResponse(
		game.ID,
		game.Name,
		game.ReleaseYear,
		game.TeamID,
		game.Universe,
		game.CreatedAt,
		game.UpdatedAt,
		game.HeroID,
		game.Consoles,
	), nil
}
func (s *service) Delete(ctx context.Context, id string) error {
	if err := s.delete(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err)
		return err
	}
	return nil
}

func (s *service) delete(ctx context.Context, id string) error {
	if err := s.cache.DeleteGame(ctx, id); err != nil {
		return exception.New(fmt.Sprintf("deleteGame\n%s", err.Error()), err)
	}

	if err := s.repository.DeleteGameByID(ctx, id); err != nil {
		return exception.New(fmt.Sprintf("deleteGameByID\n%s", err.Error()), err)
	}
	return nil
}
