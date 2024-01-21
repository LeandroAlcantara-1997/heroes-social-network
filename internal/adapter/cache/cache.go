package cache

import (
	"context"

	game "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/model"
	hero "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/model"
	team "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/model"
)

//go:generate mockgen -destination ../../mock/cache_mock.go -package=mock -source=cache.go
type Cache interface {
	HeroCache
	TeamCache
	GameCache
}

type HeroCache interface {
	SetHero(ctx context.Context, hero *hero.Hero) (err error)
	GetHero(ctx context.Context, key string) (hero *hero.Hero, err error)
	DeleteHero(ctx context.Context, key string) (err error)
}

type TeamCache interface {
	SetTeam(ctx context.Context, team *team.Team, key string) (err error)
	GetTeam(ctx context.Context, key string) (team *team.Team, err error)
	DeleteTeam(ctx context.Context, key string) (err error)
}

type GameCache interface {
	SetGame(ctx context.Context, game *game.Game) (err error)
	GetGame(ctx context.Context, key string) (game *game.Game, err error)
	DeleteGame(ctx context.Context, key string) (err error)
}
