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
	if hero.Team != nil {
		r.checkIfExistsTeam(ctx, *hero.Team)
	}
	tx, err := r.client.Begin(ctx)
	if err != nil {
		return err
	}

	query := `INSERT INTO character (id, character_name, civil_name, heroe, universe, fk_team)
		VALUES ($1, $2, $3, $4, $5, $6);`
	tag, err := tx.Exec(ctx, query,
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
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("cannot be insert hero")
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r *reposiotry) checkIfExistsTeam(ctx context.Context, id string) (bool, error) {
	var (
		query = `SELECT COUNT(id) FROM team
	WHERE id = $1;`
		count int
	)
	out, err := r.client.Query(ctx, query, id)
	if err != nil {
		return false, err
	}

	for out.Next() {
		if err := out.Scan(&count); err != nil {
			return false, err
		}
	}
	return count != 0, nil
}
