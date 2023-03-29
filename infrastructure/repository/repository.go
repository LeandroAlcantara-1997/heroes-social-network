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

func (r *reposiotry) CreateHero(ctx context.Context, hero *model.Hero) (err error) {
	if hero.Team != nil {
		if exists, err := r.checkIfExistsTeam(ctx, *hero.Team); err != nil || !exists {
			if err != nil {
				return err
			}
			if !exists {
				return fmt.Errorf("team do not exists")
			}
		}
	}

	tx, err := r.client.Begin(ctx)
	if err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return
		}
		return
	}

	var query = `INSERT INTO character (id, character_name, civil_name, hero, universe, fk_team)
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
		return
	}

	if tag.RowsAffected() == 0 {
		if err = tx.Rollback(ctx); err != nil {
			return
		}
		return fmt.Errorf("cannot be insert hero")
	}

	if err = tx.Commit(ctx); err != nil {
		return
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

func (r *reposiotry) UpdateHero(ctx context.Context, hero *model.Hero) (err error) {
	var query = `UPDATE heroes
		SET character_name = $1, civil_name = $2, hero = $3, universe = $4
		WHERE id = $5;`

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return
	}

	tag, err := tx.Exec(ctx,
		query,
		hero.HeroName, hero.CivilName, hero.Hero, hero.Universe, hero.Id)
	if err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return
		}
		return
	}

	if tag.RowsAffected() == 0 {
		if err = tx.Rollback(ctx); err != nil {
			return
		}
		return fmt.Errorf("unable to update hero")
	}
	if err = tx.Commit(ctx); err != nil {
		return
	}
	return
}

func (r *reposiotry) GetHeroByID(ctx context.Context, id string) (*model.Hero, error) {
	var (
		query = `SELECT id, character_name, civil_name, hero, universe
	FROM character
	WHERE id = $1;`
		hero = &model.Hero{}
	)

	row := r.client.QueryRow(ctx, query, id)
	if err := row.Scan(&hero.Id, &hero.HeroName, &hero.CivilName,
		&hero.Hero, &hero.Universe); err != nil {
		return nil, err
	}

	return hero, nil
}

func (r *reposiotry) DeleteHeroByID(ctx context.Context, id string) (err error) {
	var query = `DELETE FROM character 
	WHERE id = $1;`

	tag, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("unable to delete hero")
	}
	return
}
