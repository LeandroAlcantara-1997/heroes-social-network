package repository

import (
	"context"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
)

//go:generate mockgen -destination ../../../mock/repository_mock.go -package=mock -source=repository.go
type Repository interface {
	HeroRepository
	TeamRepository
}

type HeroRepository interface {
	CreateHero(ctx context.Context, hero *model.Hero) (*model.Hero, error)
	UpdateHero(ctx context.Context, hero *model.Hero) (err error)
	GetHeroByID(ctx context.Context, id string) (*model.Hero, error)
	DeleteHeroByID(ctx context.Context, id string) (err error)
}

type TeamRepository interface {
	CreateTeam(ctx context.Context, team *model.Team) (*model.Team, error)
	GetTeamByID(ctx context.Context, id string) (*model.Team, error)
	DeleteTeamByID(ctx context.Context, id string) (err error)
}
