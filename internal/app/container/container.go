package container

import (
	"context"
	"fmt"

	"github.com/LeandroAlcantara-1997/heroes-social-network/config/env"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/cache"
	log "github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/repository"
	ability "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/service"
	console "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/service"
	game "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/service"
	hero "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/service"
	team "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/otel"
	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
	redisOtel "github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

type Container struct {
	HeroUseCase    hero.Hero
	TeamUseCase    team.Team
	GameUseCase    game.Game
	ConsoleUseCase console.Console
	AbilityUseCase ability.Ability
}

type components struct {
	pgxClient    *pgxpool.Pool
	splunkClient *log.Splunk
	redisClient  *redis.Client
}

func New() (context.Context, *Container, error) {
	env.LoadEnv()
	ctx := context.Background()
	otel.New(env.Env.ServiceName, env.Env.Environment).TraceProvider(ctx)

	cmp, err := setupComponents(ctx)
	if err != nil {
		return ctx, nil, err
	}
	heroService := hero.New(
		repository.New(cmp.pgxClient),
		cache.New(cmp.redisClient),
		cmp.splunkClient,
	)

	teamService := team.New(
		repository.New(cmp.pgxClient),
		cache.New(cmp.redisClient),
		cmp.splunkClient,
	)

	gameService := game.New(
		repository.New(cmp.pgxClient),
		cache.New(cmp.redisClient),
		cmp.splunkClient,
	)

	consoleService := console.New(
		repository.New(cmp.pgxClient),
		cmp.splunkClient,
	)

	abilityService := ability.New(
		repository.New(cmp.pgxClient),
		cache.New(cmp.redisClient),
		cmp.splunkClient,
	)

	return ctx, &Container{
		HeroUseCase:    heroService,
		TeamUseCase:    teamService,
		GameUseCase:    gameService,
		ConsoleUseCase: consoleService,
		AbilityUseCase: abilityService,
	}, nil
}

func setupComponents(ctx context.Context) (*components, error) {
	pgxConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.Env.DBUser,
		env.Env.DBPassword,
		env.Env.DBHost,
		env.Env.DBPort,
		env.Env.DBName,
	))
	if err != nil {
		return nil, err
	}
	pgxConfig.ConnConfig.Tracer = otelpgx.NewTracer()
	pgxClient, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, err
	}

	splunkClient := log.New(env.Env.SplunkHost, env.Env.SplunkToken, false)

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", env.Env.CacheHost, env.Env.CachePort),
			Password: env.Env.CachePassword,
			DB:       0,
		},
	)
	redisOtel.InstrumentTracing(redisClient)

	return &components{
		pgxClient:    pgxClient,
		splunkClient: splunkClient,
		redisClient:  redisClient,
	}, nil
}
