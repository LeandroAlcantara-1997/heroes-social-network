package repository

import (
	"context"
	"errors"
	"fmt"

	game "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/jackc/pgx/v5"
)

func (r *repository) CreateGame(ctx context.Context, game *game.Game) (err error) {
	var query = `INSERT INTO game (id, name, release_year,universe, created_at)
	VALUES ($1, $2, $3, $4, $5);`

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return
	}

	if _, err = tx.Exec(ctx, query, game.ID, game.Name,
		game.ReleaseYear, game.Universe, game.CreatedAt); err != nil {
		return
	}

	if game.TeamID != nil {
		if err := r.createRelationShipTeamGame(ctx, game.ID, *game.TeamID, tx); err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return rollbackErr
			}

			return exception.New(fmt.Sprintf("createRelationShipTeamGame -> %s", err.Error()), exception.ErrTeamNotFound)
		}
	}

	if game.HeroID != nil {
		if err := r.createRelationShipHeroGame(ctx, game.ID, game.HeroID, tx); err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return rollbackErr
			}

			return exception.New(fmt.Sprintf("createRelationShipHeroGame -> %s", err.Error()), exception.ErrHeroNotFound)
		}
	}

	if err := r.createRelationShipGameConsole(ctx, tx, game.ID, game.Consoles); err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return rollbackErr
		}

		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return
	}

	return nil
}

func (r *repository) GetGameByID(ctx context.Context, id string) (*game.Game, error) {
	var (
		query = `SELECT id, name, release_year, universe, created_at, updated_at
	FROM game
	WHERE id = $1;`
		game = &game.Game{}
	)

	row := r.client.QueryRow(ctx, query, id)
	if err := row.Scan(&game.ID, &game.Name, &game.ReleaseYear, &game.Universe, &game.CreatedAt, &game.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, exception.ErrGameNotFound
		}
		return nil, err
	}

	return game, nil
}

func (r *repository) DeleteGameByID(ctx context.Context, id string) (err error) {
	tx, err := r.client.Begin(ctx)
	if err != nil {
		return
	}
	if err = r.deleteRelationShipHeroGame(ctx, &id, nil, tx); err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return
		}
		return
	}
	var query = `DELETE FROM game 
	WHERE id = $1;`

	tag, err := tx.Exec(ctx, query, id)
	if err != nil {
		return
	}

	if tag.RowsAffected() == 0 {
		return exception.ErrGameNotFound
	}

	if err = tx.Commit(ctx); err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return
		}
	}
	return
}

func (r *repository) UpdateGame(ctx context.Context, game *game.Game) (err error) {
	var query = `UPDATE heroes
		SET name = $1, teamId = $2, heroId = $3, universe = $4, release_year = $5, updated_at = $6
		WHERE id = $6;`

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return
	}

	tag, err := tx.Exec(ctx,
		query,
		game.Name, game.TeamID, game.HeroID, game.Universe, game.ReleaseYear, game.UpdatedAt)
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
		return fmt.Errorf("unable to update game")
	}
	if err = tx.Commit(ctx); err != nil {
		return
	}
	return
}
