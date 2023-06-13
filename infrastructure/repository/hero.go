package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
	"github.com/jackc/pgx/v5"
)

func (r *repository) CreateHero(ctx context.Context, hero *model.Hero) (*model.Hero, error) {
	if heroGetByName, _ := r.GetHeroByName(ctx, hero.HeroName); heroGetByName != nil {
		return heroGetByName, nil
	}
	if hero.Team != nil {
		if exists, err := r.checkIfExistsTeam(ctx, *hero.Team); err != nil || !exists {
			if err != nil {
				return nil, err
			}
			if !exists {
				return nil, errors.New("team do not exists")
			}
		}
	}

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return nil, err
	}

	var query = `INSERT INTO character (id, character_name, civil_name, 
		hero, universe, created_at, fk_team)
		VALUES ($1, $2, $3, $4, $5, $6, $7);`

	tag, err := tx.Exec(ctx, query,
		hero.ID,
		hero.HeroName,
		hero.CivilName,
		hero.Hero,
		hero.Universe,
		hero.CreatedAt,
		hero.Team,
	)
	if err != nil {
		return nil, err
	}

	if tag.RowsAffected() == 0 {
		if err = tx.Rollback(ctx); err != nil {
			return nil, err
		}
		return nil, errors.New("cannot be insert hero")
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return hero, nil
}

func (r *repository) UpdateHero(ctx context.Context, hero *model.Hero) (err error) {
	var query = `UPDATE heroes
		SET character_name = $1, civil_name = $2, hero = $3, universe = $4, updated_at = $5
		WHERE id = $6;`

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return
	}

	tag, err := tx.Exec(ctx,
		query,
		hero.HeroName, hero.CivilName, hero.Hero, hero.Universe, hero.ID)
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

func (r *repository) GetHeroByID(ctx context.Context, id string) (*model.Hero, error) {
	var (
		query = `SELECT id, character_name, civil_name, hero, universe
	FROM character
	WHERE id = $1;`
		hero = &model.Hero{}
	)

	row := r.client.QueryRow(ctx, query, id)
	if err := row.Scan(&hero.ID, &hero.HeroName, &hero.CivilName,
		&hero.Hero, &hero.Universe); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, exception.ErrHeroNotFound
		}
		return nil, err
	}

	return hero, nil
}

func (r *repository) DeleteHeroByID(ctx context.Context, id string) (err error) {
	var query = `DELETE FROM character 
	WHERE id = $1;`

	tag, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return
	}

	if tag.RowsAffected() == 0 {
		return exception.ErrHeroNotFound
	}
	return
}

func (r *repository) GetHeroByName(ctx context.Context, name string) (*model.Hero, error) {
	var (
		query = `SELECT id, character_name, civil_name, hero, 
		universe, created_at, updated_at
		FROM character
		WHERE character_name IN ($1);`
		hero = &model.Hero{}
	)

	row := r.client.QueryRow(ctx, query, name)
	if err := row.Scan(&hero.ID, &hero.HeroName, &hero.CivilName,
		&hero.Hero, &hero.Universe, &hero.CreatedAt, &hero.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, exception.ErrHeroNotFound
		}
		return nil, err
	}

	return hero, nil
}
