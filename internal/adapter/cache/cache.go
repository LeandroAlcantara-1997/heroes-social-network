package cache

import (
	"context"

	hero "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/model"
	team "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/model"
)

//go:generate mockgen -destination ../../mock/cache_mock.go -package=mock -source=cache.go
type Cache interface {
	HeroCache
	TeamCache
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
