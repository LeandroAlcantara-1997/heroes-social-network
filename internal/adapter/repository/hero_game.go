package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *repository) createRelationShipHeroGame(ctx context.Context, gameID, heroID string, tx pgx.Tx) (err error) {
	var query = `INSERT INTO character_game(character_fk, game_fk)
	VALUES ($1, $2);`

	_, err = tx.Exec(ctx, query, heroID, gameID)
	return err
}
