package repository

import (
	"context"
	"fmt"
)

func (r *repository) getAbilitiesByHeroID(ctx context.Context, id string) ([]string, error) {
	var (
		query = `SELECT fk_ability FROM character_ability 
	WHERE fk_character = $1;`
		ids = make([]string, 0)
	)
	rows, err := r.client.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("query\n%w", err)
	}

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan\n%w", err)
		}
		ids = append(ids, id)
	}

	defer rows.Close()

	return ids, nil
}

func (r *repository) reateRelationShipAbilityHero(ctx context.Context, abilityID, heroID string) error {
	var query = `INSERT INTO character_ability (fk_character, fk_ability)
		VALUES ($1, $2);`

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin\n%w", err)
	}

	_, err = tx.Exec(ctx, query, heroID, abilityID)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return fmt.Errorf("rollback\n%w", err)
		}
		return fmt.Errorf("exec\n%w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("\n%w", err)
	}
	return nil
}
