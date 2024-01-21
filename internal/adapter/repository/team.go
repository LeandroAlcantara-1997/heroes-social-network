package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
)

func (r *repository) CreateTeam(ctx context.Context, team *model.Team) error {
	var query = `INSERT INTO team (id, name, universe, created_at)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (name)
	DO NOTHING;`

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return err
	}

	tag, err := tx.Exec(ctx, query, team.ID,
		team.Name, team.Universe, team.CreatedAt)
	if err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return err
		}
		return errors.New("cannot be insert team")
	}

	if tag.RowsAffected() == 0 {
		if err = tx.Rollback(ctx); err != nil {
			return err
		}
		return exception.ErrTeamAlredyExists
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (r *repository) GetTeamByID(ctx context.Context, id string) (*model.Team, error) {
	var (
		query = `SELECT id, name, universe, created_at, updated_at FROM team
		WHERE id = $1;`
		team = &model.Team{}
	)

	row := r.client.QueryRow(ctx, query, id)
	if err := row.Scan(&team.ID, &team.Name, &team.Universe,
		&team.CreatedAt, &team.UpdatedAt); err != nil {
		return nil, err
	}

	return team, nil

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

func (r *repository) DeleteTeamByID(ctx context.Context, id string) (err error) {
	var query = `DELETE FROM team 
	WHERE id = $1;`

	tag, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return
	}

	if tag.RowsAffected() == 0 {
		return exception.ErrTeamNotFound
	}
	return
}

func (r *repository) GetTeamByName(ctx context.Context, name string) (*model.Team, error) {
	var (
		query = `SELECT id, name, universe, created_at, updated_at FROM team
		WHERE name IN ($1);`
		team = &model.Team{}
	)

	row := r.client.QueryRow(ctx, query, name)
	if err := row.Scan(&team.ID, &team.Name, &team.Universe,
		&team.CreatedAt, &team.UpdatedAt); err != nil {
		return nil, err
	}

	return team, nil
}

func (r *repository) UpdateTeam(ctx context.Context, team *model.Team) (err error) {
	var query = `UPDATE team
		SET name = $1, universe = $2, updated_at = $3
		WHERE id = '$4';`
	tx, err := r.client.Begin(ctx)
	if err != nil {
		return
	}

	tag, err := tx.Exec(ctx,
		query,
		team.Name, team.Universe, team.UpdatedAt, team.ID)
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
		return fmt.Errorf("unable to update team")
	}
	if err = tx.Commit(ctx); err != nil {
		return
	}
	return
}
