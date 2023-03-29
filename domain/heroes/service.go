package heroes

import (
	"context"

	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/output/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/output/repository"
	"github.com/google/uuid"
)

type service struct {
	repository repository.Repository
	log        log.Log
}

func New(repository repository.Repository, log log.Log) *service {
	return &service{
		repository: repository,
		log:        log,
	}
}

func (s *service) RegisterHero(ctx context.Context, dto *input.HeroRequest) (*input.HeroResponse, error) {
	hero := model.New(uuid.NewString(), dto)
	if !model.CheckUniverse(model.Universe(hero.Universe)) {
		s.log.SendErrorLog(ctx, "invalid field")
		return nil, exception.New(exception.InvalidFieldsError)
	}

	if err := s.repository.CreateHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		return nil, exception.New(exception.InternalServerError)
	}

	return &input.HeroResponse{
		Id:        hero.Id,
		HeroName:  hero.HeroName,
		CivilName: hero.CivilName,
		Hero:      hero.Hero,
		Universe:  hero.Universe,
	}, nil
}

func (s *service) UpdateHero(ctx context.Context, id string, dto *input.HeroRequest) (*input.HeroResponse, error) {
	hero := model.New(id, dto)
	if !model.CheckUniverse(model.Universe(hero.Universe)) {
		s.log.SendErrorLog(ctx, "invalid field")
		return nil, exception.New(exception.InvalidFieldsError)
	}

	if err := s.repository.UpdateHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		return nil, err
	}

	return &input.HeroResponse{
		Id:        id,
		HeroName:  hero.HeroName,
		CivilName: hero.CivilName,
		Hero:      hero.Hero,
		Universe:  hero.Universe,
	}, nil
}

func (s *service) GetHeroByID(ctx context.Context, id string) (*input.HeroResponse, error) {
	hero, err := s.repository.GetHeroByID(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		return nil, exception.New(exception.InternalServerError)
	}
	return &input.HeroResponse{
		Id:        hero.Id,
		HeroName:  hero.HeroName,
		CivilName: hero.CivilName,
		Hero:      hero.Hero,
		Universe:  hero.Universe,
	}, nil
}

func (s *service) DeleteHeroByID(ctx context.Context, id string) (err error) {
	if err = s.repository.DeleteHeroByID(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err.Error())
		return exception.New(exception.InternalServerError)
	}
	return
}
