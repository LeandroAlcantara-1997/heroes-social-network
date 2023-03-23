package repository

import (
	"context"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
	"github.com/jackc/pgx/v5"
)

type reposiotry struct {
	client *pgx.Conn
}

func New(client *pgx.Conn) *reposiotry {
	return &reposiotry{
		client: client,
	}
}

func (r *reposiotry) CreateHero(ctx context.Context, hero model.Hero) error {
	return nil
}
