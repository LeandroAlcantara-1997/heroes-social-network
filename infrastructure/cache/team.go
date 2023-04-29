package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
)

func (c *cache) SetTeam(ctx context.Context, team *model.Team) (err error) {
	payload, err := json.Marshal(team)
	if err != nil {
		return
	}

	if cmd := c.redisClient.Set(ctx,
		getTeamKey(team.Id), payload,
		time.Duration(time.Hour*24)); cmd.Err() != nil {
		return cmd.Err()
	}

	return
}

func (c *cache) GetTeam(ctx context.Context, key string) (*model.Team, error) {
	var (
		team *model.Team
		cmd  = c.redisClient.Get(ctx, getTeamKey(key))
	)

	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	out, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(out), &team); err != nil {
		return nil, err
	}
	return team, nil
}

func (c *cache) DeleteTeam(ctx context.Context, key string) (err error) {
	cmd := c.redisClient.Del(ctx, getTeamKey(key))
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return
}

func getTeamKey(id string) string {
	return fmt.Sprintf("team:%s", id)
}
