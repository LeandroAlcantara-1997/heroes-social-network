package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	ability "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/model"
)

func (c *cache) SetAbility(ctx context.Context, ability *ability.Ability) (err error) {
	payload, err := json.Marshal(ability)
	if err != nil {
		return
	}

	if cmd := c.redisClient.Set(ctx,
		getAbilityKey(ability.ID), payload,
		time.Duration(time.Hour*24)); cmd.Err() != nil {
		return cmd.Err()
	}

	return
}

func (c *cache) GetAbility(ctx context.Context, key string) (*ability.Ability, error) {
	var (
		ability *ability.Ability
		cmd     = c.redisClient.Get(ctx, getGameKey(key))
	)

	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	out, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(out), &ability); err != nil {
		return nil, err
	}
	return ability, nil
}

func (c *cache) DeleteAbility(ctx context.Context, key string) (err error) {
	cmd := c.redisClient.Del(ctx, getGameKey(key))
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return
}

func getAbilityKey(key string) string {
	return fmt.Sprintf("ability:%s", key)
}
