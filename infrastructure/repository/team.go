package repository

import (
	"context"
	"errors"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
)

func (r *repository) CreateTeam(ctx context.Context, team *model.Team) (err error) {
	var query = `INSERT INTO team (id, name, universe, created_at)
	VALUES ($1, $2, $3, $4);`

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return
	}

	tag, err := tx.Exec(ctx, query, team.Id,
		team.Name, team.Universe, team.CreatedAt)
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
		return errors.New("cannot be insert team")
	}

	if err = tx.Commit(ctx); err != nil {
		return
	}
	return
}

func (r *repository) checkIfExistsTeam(ctx context.Context, id string) (bool, error) {
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
