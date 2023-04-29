package cache

import (
	"context"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
)

//go:generate mockgen -destination ../../../mock/cache_mock.go -package=mock -source=cache.go
type Cache interface {
	HeroCache
	TeamCache
}

type HeroCache interface {
	SetHero(ctx context.Context, hero *model.Hero) (err error)
	GetHero(ctx context.Context, key string) (hero *model.Hero, err error)
	DeleteHero(ctx context.Context, key string) (err error)
}

type TeamCache interface {
	SetTeam(ctx context.Context, team *model.Team) (err error)
	GetTeam(ctx context.Context, key string) (team *model.Team, err error)
	DeleteTeam(ctx context.Context, key string) (err error)
}
