package repository

import (
	"context"

	console "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/model"
	"github.com/jackc/pgx/v5"
)

func (r *repository) createRelationShipGameConsole(ctx context.Context, tx pgx.Tx, gameID string, consoles []console.Console) error {
	var query = `INSERT INTO console_game(fk_console, fk_game)
	VALUES ($1, $2);`
	for c := range consoles {
		if _, err := tx.Exec(ctx, query, consoles[c], gameID); err != nil {
			return err
		}
	}

	return nil
}
