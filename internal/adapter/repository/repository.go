package repository

import (
	"context"

	heroes "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/model"
	teams "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/model"
)

//go:generate mockgen -destination ../../mock/repository_mock.go -package=mock -source=repository.go
type Repository interface {
	HeroRepository
	TeamRepository
}

type HeroRepository interface {
	CreateHero(ctx context.Context, hero *heroes.Hero) error
	UpdateHero(ctx context.Context, hero *heroes.Hero) (err error)
	GetHeroByID(ctx context.Context, id string) (*heroes.Hero, error)
	DeleteHeroByID(ctx context.Context, id string) (err error)
}

type TeamRepository interface {
	CreateTeam(ctx context.Context, team *teams.Team) error
	GetTeamByID(ctx context.Context, id string) (*teams.Team, error)
	GetTeamByName(ctx context.Context, name string) (*teams.Team, error)
	DeleteTeamByID(ctx context.Context, id string) (err error)
	UpdateTeam(ctx context.Context, team *teams.Team) (err error)
}
