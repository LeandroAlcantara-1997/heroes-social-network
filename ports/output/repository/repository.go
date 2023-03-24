package repository

import (
	"context"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
)

//go:generate mockgen -destination ../../../mock/repository_mock.go -package=mock -source=repository.go
type Repository interface {
	CreateHero(ctx context.Context, hero *model.Hero) (err error)
	// GetTeamById(ctx context.Context, id string)
}
