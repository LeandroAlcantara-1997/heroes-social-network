package repository

import (
	"context"

	console "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/model"
)

func (r *repository) CreateConsoles(ctx context.Context, consoles []console.Console) error {
	var query = `INSERT INTO console (name)
	VALUES ($1);`
	tx, err := r.client.Begin(ctx)
	if err != nil {
		return err
	}

	for c := range consoles {
		if _, err := tx.Exec(ctx, query, consoles[c]); err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return rollbackErr
			}
			return err
		}
	}

	return tx.Commit(ctx)

}

func (r *repository) GetConsoles(ctx context.Context) ([]console.Console, error) {
	var (
		query = `SELECT *
	FROM console;`
		consoles []console.Console
		console  console.Console
	)

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&console); err != nil {
			return nil, err
		}
		consoles = append(consoles, console)
	}

	defer rows.Close()
	return consoles, nil
}
