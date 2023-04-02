package cache

import (
	"context"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
)

//go:generate mockgen -destination ../../../mock/cache_mock.go -package=mock -source=cache.go
type Cache interface {
	Set(ctx context.Context, hero *model.Hero) (err error)
	Get(ctx context.Context, key string) (hero *model.Hero, err error)
	Delete(ctx context.Context, key string) (err error)
}
