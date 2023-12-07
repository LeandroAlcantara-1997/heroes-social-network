package hero

import (
	"context"
	"errors"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/cache"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/repository"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/model"
	"github.com/google/uuid"
)

//go:generate mockgen -destination ../../../mock/hero_mock.go -package=mock -source=service.go
type Hero interface {
	RegisterHero(ctx context.Context, request *dto.HeroRequest) (*dto.HeroResponse, error)
	UpdateHero(ctx context.Context, id string, request *dto.HeroRequest) error
	GetHeroByID(ctx context.Context, id string) (*dto.HeroResponse, error)
	DeleteHeroByID(ctx context.Context, id string) (err error)
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

func (s *service) RegisterHero(ctx context.Context, request *dto.HeroRequest) (*dto.HeroResponse, error) {

	hero := model.NewHero(uuid.NewString(), request)
	if err := s.repository.CreateHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}

	if err := s.cache.SetHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, exception.ErrInternalServer
	}

	return dto.NewHeroResponse(hero.ID, hero.HeroName, hero.CivilName, hero.Universe,
		hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
}

func (s *service) UpdateHero(ctx context.Context, id string, request *dto.HeroRequest) error {
	hero := model.NewHero(id, request)
	hero.UpdatedAt = util.GerPointer(time.Now().UTC())
	if err := s.repository.UpdateHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err)
		if err := s.cache.DeleteHero(ctx, hero.ID); err != nil {
			s.log.SendErrorLog(ctx, err)
			return exception.ErrInternalServer
		}
		return err
	}

	if err := s.cache.SetHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err)
		if err := s.cache.DeleteHero(ctx, hero.ID); err != nil {
			s.log.SendErrorLog(ctx, err)
		}
	}

	return nil
}

func (s *service) GetHeroByID(ctx context.Context, id string) (*dto.HeroResponse, error) {
	hero, err := s.cache.GetHero(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		hero, err = s.repository.GetHeroByID(ctx, id)
		if err != nil {
			s.log.SendErrorLog(ctx, err)
			if errors.Is(err, exception.ErrHeroNotFound) {
				return nil, err
			}
			return nil, exception.ErrInternalServer
		}
		return dto.NewHeroResponse(hero.ID, hero.HeroName, hero.CivilName,
			hero.Universe, hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
	}

	return dto.NewHeroResponse(hero.ID, hero.HeroName, hero.CivilName,
		hero.Universe, hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
}

func (s *service) DeleteHeroByID(ctx context.Context, id string) (err error) {
	if err = s.cache.DeleteHero(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err)
		return exception.ErrInternalServer
	}

	if err = s.repository.DeleteHeroByID(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err)
		if errors.Is(err, exception.ErrHeroNotFound) {
			return
		}
		return exception.ErrInternalServer
	}
	return
}
