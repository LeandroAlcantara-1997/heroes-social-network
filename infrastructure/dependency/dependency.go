package dependency

import (
	"context"
	"fmt"

	"github.com/LeandroAlcantara-1997/heroes-social-network/domain/heroes"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/cache"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/config"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/repository"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/splunk"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type dependency struct {
	HeroUseCase input.Hero
}

type components struct {
	pgxClient    *pgx.Conn
	splunkClient *splunk.Splunk
	redisClient  *redis.Client
}

func LoadDependency(ctx context.Context) (*dependency, error) {
	cmp, err := loadExternalTools(ctx)
	if err != nil {
		return nil, err
	}
	heroService := heroes.New(
		repository.New(cmp.pgxClient),
		cache.New(cmp.redisClient),
		cmp.splunkClient,
	)

	return &dependency{
		HeroUseCase: heroService,
	}, nil
}

func loadExternalTools(ctx context.Context) (*components, error) {
	pgxClient, err := pgx.Connect(
		ctx,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.Env.DbUser,
			config.Env.DbPassword,
			config.Env.DbHost,
			config.Env.DbPort,
			config.Env.DbName,
		),
	)
	if err != nil {
		return nil, err
	}

	splunkClient := splunk.New(config.Env.SplunkHost, false)

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.Env.CacheHost, config.Env.CachePort),
			Password: config.Env.CachePassword,
			DB:       0,
		},
	)

	return &components{
		pgxClient:    pgxClient,
		splunkClient: splunkClient,
		redisClient:  redisClient,
	}, nil
}
