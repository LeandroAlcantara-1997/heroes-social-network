package heroes

import (
	"context"
	"errors"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
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
	hero := model.New(uuid.NewString(), dto)
	if !model.CheckUniverse(model.Universe(hero.Universe)) {
		s.log.SendErrorLog(ctx, "invalid field")
		return nil, exception.ErrInvalidFields
	}

	if err := s.repository.CreateHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		if err := s.cache.Delete(ctx, hero.Id); err != nil {
			s.log.SendErrorLog(ctx, err.Error())
			return nil, exception.ErrInternalServer
		}
		return nil, exception.ErrInternalServer
	}

	if err := s.cache.Set(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		return nil, exception.ErrInternalServer
	}

	return input.NewHeroResponse(hero.Id, hero.HeroName, hero.CivilName, hero.Universe,
		hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
}

func (s *service) UpdateHero(ctx context.Context, id string, dto *input.HeroRequest) (*input.HeroResponse, error) {
	hero := model.New(id, dto)
	if !model.CheckUniverse(model.Universe(hero.Universe)) {
		s.log.SendErrorLog(ctx, "invalid field")
		return nil, exception.ErrInvalidFields
	}
	hero.UpdatedAt = gerPointer(time.Now().UTC())
	if err := s.repository.UpdateHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		if err := s.cache.Delete(ctx, hero.Id); err != nil {
			s.log.SendErrorLog(ctx, err.Error())
			return nil, exception.ErrInternalServer
		}
		return nil, err
	}

	if err := s.cache.Set(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		return nil, exception.ErrInternalServer
	}

	return input.NewHeroResponse(id, hero.HeroName, hero.CivilName, hero.Universe,
		hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
}

func (s *service) GetHeroByID(ctx context.Context, id string) (*input.HeroResponse, error) {
	hero, err := s.cache.Get(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		hero, err = s.repository.GetHeroByID(ctx, id)
		if err != nil {
			s.log.SendErrorLog(ctx, err.Error())
			if errors.Is(err, exception.ErrHeroNotFound) {
				return nil, err
			}
			return nil, exception.ErrInternalServer
		}
		return input.NewHeroResponse(hero.Id, hero.HeroName, hero.CivilName,
			hero.Universe, hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
	}

	return input.NewHeroResponse(hero.Id, hero.HeroName, hero.CivilName,
		hero.Universe, hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
}

func (s *service) DeleteHeroByID(ctx context.Context, id string) (err error) {
	if err = s.cache.Delete(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		return exception.ErrInternalServer
	}

	if err = s.repository.DeleteHeroByID(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
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
