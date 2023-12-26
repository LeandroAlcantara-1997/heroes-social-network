package repository

import (
	"context"

	game "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/model"
)

func (r *repository) CreateGame(ctx context.Context, game *game.Game) (err error) {
	var query = `INSERT INTO game (id, name, release_year,universe)
	VALUES ($1, $2, $3, $4);`

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return
	}
	_, err = tx.Exec(ctx, query, game.ID, game.Name,
		game.ReleaseYear, game.Universe)
	if err != nil {
		return
	}
	if game.TeamID != nil {
		if err := r.createRelationShipTeamGame(ctx, game.ID, *game.TeamID, tx); err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return rollbackErr
			}

			return err
		}
	}

	if game.HeroID != nil {
		if err := r.createRelationShipHeroGame(ctx, game.ID, *game.TeamID, tx); err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return rollbackErr
			}

			return err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return
	}

	return nil
}
