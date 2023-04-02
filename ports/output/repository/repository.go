package repository

import (
	"context"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
)

//go:generate mockgen -destination ../../../mock/repository_mock.go -package=mock -source=repository.go
type Repository interface {
	HeroRepository
}

type HeroRepository interface {
	CreateHero(ctx context.Context, hero *model.Hero) (err error)
	UpdateHero(ctx context.Context, hero *model.Hero) (err error)
	GetHeroByID(ctx context.Context, id string) (*model.Hero, error)
	DeleteHeroByID(ctx context.Context, id string) (err error)
}
