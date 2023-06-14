package repository

import (
	"context"
	"errors"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
)

func (r *repository) CreateTeam(ctx context.Context, team *model.Team) (*model.Team, error) {
	if teamGetByName, _ := r.GetTeamByName(ctx, team.Name); teamGetByName != nil {
		return teamGetByName, nil
	}
	var query = `INSERT INTO team (id, name, universe, created_at)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (name)
	DO NOTHING;`

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return nil, err
	}

	tag, err := tx.Exec(ctx, query, team.ID,
		team.Name, team.Universe, team.CreatedAt)
	if err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return nil, err
		}
		return nil, errors.New("cannot be insert team")
	}

	if tag.RowsAffected() == 0 {
		if err = tx.Rollback(ctx); err != nil {
			return nil, err
		}
		return nil, errors.New("cannot be insert team")
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}
	return team, nil
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
