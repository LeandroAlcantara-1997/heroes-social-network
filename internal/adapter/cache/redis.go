package cache

import "github.com/redis/go-redis/v9"

type cache struct {
	redisClient *redis.Client
}

func New(client *redis.Client) *cache {
	return &cache{
		redisClient: client,
	}
}
