package heroes

import (
	"context"

	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/output/repository"
	"github.com/google/uuid"
)

type service struct {
	repository repository.Repository
}

func New(repository repository.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) RegisterHero(ctx context.Context, heroDto *input.HeroRequest) (*input.HeroResponse, error) {
	heroModel := model.New(uuid.NewString(), *heroDto)
	if !model.CheckUniverse(model.Universe(heroModel.Universe)) {
		return nil, exception.New(exception.InvalidFieldsError)
	}

	if err := s.repository.CreateHero(ctx, heroModel); err != nil {
		return nil, exception.New(exception.InternalServerError)
	}

	return &input.HeroResponse{
		Id:        heroModel.Id,
		HeroName:  heroModel.HeroName,
		CivilName: heroDto.CivilName,
		Hero:      heroModel.Hero,
		Universe:  heroModel.Universe,
	}, nil
}
