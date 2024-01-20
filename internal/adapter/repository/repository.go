package repository

import (
	"context"

	console "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/model"
	game "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/model"
	hero "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/model"
	team "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/model"
)

//go:generate mockgen -destination ../../mock/repository_mock.go -package=mock -mock_names=Repository=RepositoryMock -source=repository.go
type Repository interface {
	HeroRepository
	TeamRepository
	GameRepository
}

type HeroRepository interface {
	CreateHero(ctx context.Context, hero *hero.Hero) error
	UpdateHero(ctx context.Context, hero *hero.Hero) (err error)
	GetHeroByID(ctx context.Context, id string) (*hero.Hero, error)
	DeleteHeroByID(ctx context.Context, id string) (err error)
}

type TeamRepository interface {
	CreateTeam(ctx context.Context, team *team.Team) error
	GetTeamByID(ctx context.Context, id string) (*team.Team, error)
	GetTeamByName(ctx context.Context, name string) (*team.Team, error)
	DeleteTeamByID(ctx context.Context, id string) (err error)
	UpdateTeam(ctx context.Context, team *team.Team) (err error)
}

type GameRepository interface {
	CreateGame(ctx context.Context, game *game.Game) error
	UpdateGame(ctx context.Context, game *game.Game) (err error)
	GetGameByID(ctx context.Context, id string) (*game.Game, error)
	DeleteGameByID(ctx context.Context, id string) error
}

type ConsoleRepository interface {
	CreateConsoles(ctx context.Context, consoles []console.Console) error
	GetConsoles(ctx context.Context) ([]console.Console, error)
}
