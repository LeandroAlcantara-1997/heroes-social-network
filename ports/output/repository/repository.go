package repository

import (
	"context"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
)

type Repository interface {
	CreateHero(ctx context.Context, hero model.Hero) (err error)
}
