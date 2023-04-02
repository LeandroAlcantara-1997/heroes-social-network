package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
	"github.com/redis/go-redis/v9"
)

type cache struct {
	redisClient *redis.Client
}

func New(client *redis.Client) *cache {
	return &cache{
		redisClient: client,
	}
}

func (c *cache) Set(ctx context.Context, hero *model.Hero) (err error) {
	payload, err := json.Marshal(hero)
	if err != nil {
		return
	}

	if cmd := c.redisClient.Set(ctx,
		getKey(hero.Id), payload,
		time.Duration(time.Hour*24)); cmd.Err() != nil {
		return cmd.Err()
	}

	return
}

func (c *cache) Get(ctx context.Context, key string) (*model.Hero, error) {
	var (
		hero *model.Hero
		cmd  = c.redisClient.Get(ctx, getKey(key))
	)

	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	out, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(out), &hero); err != nil {
		return nil, err
	}
	return hero, nil
}

func (c *cache) Delete(ctx context.Context, key string) (err error) {
	cmd := c.redisClient.Del(ctx, getKey(key))
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return
}

func getKey(id string) string {
	return fmt.Sprintf("hero:%s", id)
}
