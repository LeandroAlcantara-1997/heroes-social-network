package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *repository) createRelationShipTeamGame(ctx context.Context, gameID, teamID string, tx pgx.Tx) (err error) {
	var query = `INSERT INTO team_game (fk_team, fk_game)
	VALUES ($1, $2);`

	_, err = tx.Exec(ctx, query, teamID, gameID)
	return err
}
