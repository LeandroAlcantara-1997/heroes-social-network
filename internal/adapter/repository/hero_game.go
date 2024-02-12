package repository

import (
	"context"
	"fmt"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/jackc/pgx/v5"
)

func (r *repository) createRelationShipHeroGame(ctx context.Context, gameID string, heroID []string, tx pgx.Tx) (err error) {
	var query = `INSERT INTO character_game(fk_character, fk_game)
	VALUES ($1, $2);`
	for h := range heroID {
		if _, err = tx.Exec(ctx, query, heroID[h], gameID); err != nil {
			return fmt.Errorf("exec\n%s", err)
		}
	}
	return nil
}

func (r *repository) deleteRelationShipHeroGame(ctx context.Context, gameID, heroID *string, tx pgx.Tx) (err error) {
	var (
		query = `DELETE FROM character_game 
	WHERE `
		id string
	)
	if gameID != nil {
		query += `fk_game = $1;`
		id = *gameID
	}

	if heroID != nil {
		query += `fk_character = $1;`
	}

	tag, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return
	}

	if tag.RowsAffected() == 0 {
		return exception.ErrGameNotFound
	}
	return
}
