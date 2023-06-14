package dependency

import (
	"context"
	"fmt"

	"github.com/LeandroAlcantara-1997/heroes-social-network/domain/heroes"
	"github.com/LeandroAlcantara-1997/heroes-social-network/domain/teams"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/cache"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/config"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/repository"
	log "github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/splunk"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/hero"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type dependency struct {
	HeroUseCase hero.Hero
	TeamUseCase team.Team
}

type components struct {
	pgxClient    *pgx.Conn
	splunkClient *log.Splunk
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

	teamService := teams.New(
		repository.New(cmp.pgxClient),
		cache.New(cmp.redisClient),
		cmp.splunkClient,
	)

	return &dependency{
		HeroUseCase: heroService,
		TeamUseCase: teamService,
	}, nil
}

func loadExternalTools(ctx context.Context) (*components, error) {
	pgxClient, err := pgx.Connect(
		ctx,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.Env.DBUser,
			config.Env.DBPassword,
			config.Env.DBHost,
			config.Env.DBPort,
			config.Env.DBName,
		),
	)
	if err != nil {
		return nil, err
	}

	splunkClient := log.New(config.Env.SplunkHost, false)

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
