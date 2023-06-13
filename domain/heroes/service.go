package heroes

import (
	"context"
	"errors"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
	input "github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/hero"
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

func (s *service) RegisterHero(ctx context.Context, dto *input.HeroRequest) (*input.HeroResponse, error) {
	if !model.CheckUniverse(model.Universe(dto.Universe)) {
		s.log.SendErrorLog(ctx, errors.New("invalid field"))
		return nil, exception.ErrInvalidFields
	}

	hero, err := s.repository.CreateHero(ctx, model.New(uuid.NewString(), dto))
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, exception.ErrInternalServer
	}

	if err := s.cache.SetHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, exception.ErrInternalServer
	}

	return input.NewHeroResponse(hero.ID, hero.HeroName, hero.CivilName, hero.Universe,
		hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
}

func (s *service) UpdateHero(ctx context.Context, id string, dto *input.HeroRequest) (*input.HeroResponse, error) {
	hero := model.New(id, dto)
	if !model.CheckUniverse(model.Universe(hero.Universe)) {
		s.log.SendErrorLog(ctx, errors.New("invalid field"))
		return nil, exception.ErrInvalidFields
	}
	hero.UpdatedAt = gerPointer(time.Now().UTC())
	if err := s.repository.UpdateHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err)
		if err := s.cache.DeleteHero(ctx, hero.ID); err != nil {
			s.log.SendErrorLog(ctx, err)
			return nil, exception.ErrInternalServer
		}
		return nil, err
	}

	if err := s.cache.SetHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err)
		if err := s.cache.DeleteHero(ctx, hero.ID); err != nil {
			s.log.SendErrorLog(ctx, err)
		}
	}

	return input.NewHeroResponse(id, hero.HeroName, hero.CivilName, hero.Universe,
		hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
}

func (s *service) GetHeroByID(ctx context.Context, id string) (*input.HeroResponse, error) {
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
		return input.NewHeroResponse(hero.ID, hero.HeroName, hero.CivilName,
			hero.Universe, hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
	}

	return input.NewHeroResponse(hero.ID, hero.HeroName, hero.CivilName,
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

func gerPointer[T time.Time](value T) *T {
	return &value
}
