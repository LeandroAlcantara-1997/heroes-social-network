package repository

import (
	"context"
	"errors"
	"fmt"

	ability "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/model"
)

func (r *repository) CreateAbility(ctx context.Context, ability *ability.Ability) error {
	tx, err := r.client.Begin(ctx)
	if err != nil {
		return err
	}

	var query = `INSERT INTO ability (id, description, created_at)
		VALUES ($1, $2, $3);`

	tag, err := tx.Exec(ctx, query,
		ability.ID,
		ability.Description,
		ability.CreatedAt,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		if err = tx.Rollback(ctx); err != nil {
			return err
		}
		return errors.New("cannot be insert ability")
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAbilityByID(ctx context.Context, id string) (*ability.Ability, error) {
	var (
		query = `SELECT id, description,created_at, updated_at FROM ability
	WHERE id = $1;`
		ability ability.Ability
	)

	row := r.client.QueryRow(ctx, query, id)
	if err := row.Scan(&ability.ID, &ability.Description, &ability.CreatedAt, &ability.UpdatedAt); err != nil {
		return nil, fmt.Errorf("scan\n%w", err)
	}

	return &ability, nil
}

func (r *repository) GetAbilitiesByHeroID(ctx context.Context, id string) ([]ability.Ability, error) {
	ids, err := r.getAbilitiesByHeroID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getAbilitiesByHeroID\n%w", err)
	}

	var (
		query = fmt.Sprintf(`SELECT id, description, created_at, updated_at FROM ability
	WHERE id IN (%s);`, arrayHandling(ids))
		abilities = make([]ability.Ability, len(ids))
		a         int
	)

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query\n%w", err)
	}
	if rows.Next() {
		if err := rows.Scan(&abilities[a].ID, &abilities[a].Description, &abilities[a].CreatedAt,
			&abilities[a].UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan\n%w", err)
		}
		a++
	}

	defer rows.Close()
	return abilities, nil
}

func arrayHandling(ids []string) string {
	var joinArray string
	for id := range ids {
		if id < (len(ids) - 1) {
			joinArray += fmt.Sprintf(`'%s',`, ids[id])
			continue
		}
		joinArray += fmt.Sprintf(`'%s'`, ids[id])
	}

	return joinArray
}
