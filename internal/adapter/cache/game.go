package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	game "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/model"
)

func (c *cache) SetGame(ctx context.Context, game *game.Game, key string) (err error) {
	payload, err := json.Marshal(game)
	if err != nil {
		return
	}

	if cmd := c.redisClient.Set(ctx,
		getGameKey(key), payload,
		time.Duration(time.Hour*24)); cmd.Err() != nil {
		return cmd.Err()
	}

	return
}

func (c *cache) GetGame(ctx context.Context, key string) (*game.Game, error) {
	var (
		game *game.Game
		cmd  = c.redisClient.Get(ctx, getGameKey(key))
	)

	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	out, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(out), &game); err != nil {
		return nil, err
	}
	return game, nil
}

func (c *cache) DeleteGame(ctx context.Context, key string) (err error) {
	cmd := c.redisClient.Del(ctx, getGameKey(key))
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return
}

func getGameKey(key string) string {
	return fmt.Sprintf("game:%s", key)
}
