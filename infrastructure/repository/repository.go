package repository

import (
	"context"
	"fmt"

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

func (r *reposiotry) CreateHero(ctx context.Context, hero *model.Hero) error {
	tx, err := r.client.Begin(ctx)
	if err != nil {
		return err
	}

	query := `INSERT INTO character (id, character_name, civil_name, heroe, universe, fk_team)
		VALUES ($1, $2, $3, $4, $5, $6);`
	rows, err := tx.Exec(ctx, query,
		hero.Id,
		hero.HeroName,
		hero.CivilName,
		hero.Hero,
		hero.Universe,
		hero.Team,
	)
	if err != nil {
		return err
	}
	fmt.Println(rows.RowsAffected())
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
