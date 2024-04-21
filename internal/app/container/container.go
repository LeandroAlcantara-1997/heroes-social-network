package container

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
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
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	redisOtel "github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Container struct {
	Domains    Domains
	components *components
}

type Domains struct {
	HeroUseCase    hero.Hero
	TeamUseCase    team.Team
	GameUseCase    game.Game
	ConsoleUseCase console.Console
	AbilityUseCase ability.Ability
}

type components struct {
	pgxClient   *pgxpool.Pool
	logVendor   log.Vendor
	redisClient *redis.Client
	zapLogger   *zap.Logger
}

func New() (context.Context, *Container, error) {
	env.LoadEnv()
	ctx := context.Background()
	otel.New(env.Env.ServiceName, env.Env.Environment).TraceProvider(ctx)

	cmp, err := setupComponents(ctx)
	if err != nil {
		return ctx, nil, fmt.Errorf("setupComponents -> %w", err)
	}
	heroService := hero.New(
		repository.New(cmp.pgxClient),
		cache.New(cmp.redisClient),
	)

	teamService := team.New(
		repository.New(cmp.pgxClient),
		cache.New(cmp.redisClient),
	)

	gameService := game.New(
		repository.New(cmp.pgxClient),
		cache.New(cmp.redisClient),
	)

	consoleService := console.New(
		repository.New(cmp.pgxClient),
	)

	abilityService := ability.New(
		repository.New(cmp.pgxClient),
		cache.New(cmp.redisClient),
	)

	return ctx, &Container{
		Domains: Domains{
			HeroUseCase:    heroService,
			TeamUseCase:    teamService,
			GameUseCase:    gameService,
			ConsoleUseCase: consoleService,
			AbilityUseCase: abilityService,
		},
		components: cmp,
	}, nil
}

func setupComponents(ctx context.Context) (*components, error) {
	redisClient, err := createConnectionRedis(ctx)
	if err != nil {
		return nil, fmt.Errorf("createConnectionRedis -> %w", err)
	}

	pgxClient, err := createConnectionDatabase(ctx)
	if err != nil {
		return nil, fmt.Errorf("createConnectionDatabase -> %w", err)
	}

	splunkClient := log.New(env.Env.SplunkHost, env.Env.SplunkToken, env.Env.SplunkAssync, newZapLogger(env.Env.Environment))

	return &components{
		pgxClient:   pgxClient,
		logVendor:   splunkClient,
		redisClient: redisClient,
	}, nil
}

func (c *Container) GetVendor() log.Vendor {
	return c.components.logVendor
}

func newZapLogger(environment string) *zap.Logger {
	var l, _ = zap.NewDevelopment()
	if environment != "prd" {
		l, _ = zap.NewProduction()
	}
	return l
}

func (c *Container) GetZapLogger() *zap.Logger {
	return c.components.zapLogger
}

func createConnectionDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	var conn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.Env.DBUser,
		env.Env.DBPassword,
		env.Env.DBHost,
		env.Env.DBPort,
		env.Env.DBName,
	)
	pgxConfig, err := pgxpool.ParseConfig(conn)
	if err != nil {
		return nil, fmt.Errorf("parseConfig -> %w", err)
	}
	pgxConfig.ConnConfig.Tracer = otelpgx.NewTracer()
	pgxClient, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("newWithConfig -> %w", err)
	}
	if err = pgxClient.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping -> %w", err)
	}
	sqlDB, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("open -> %w", err)
	}
	databaseDriver, err := pgx.WithInstance(sqlDB, &pgx.Config{
		DatabaseName: env.Env.DBName,
	})
	if err != nil {
		return nil, fmt.Errorf("withInstance -> %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./config/migration",
		env.Env.DBName,
		databaseDriver,
	)
	if err != nil {
		return nil, fmt.Errorf("newWithDatabaseInstance -> %w", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("up -> %w", err)
	}
	return pgxClient, nil
}

func createConnectionRedis(ctx context.Context) (*redis.Client, error) {
	redisClient := redis.NewClient(
		&redis.Options{
			Addr:         fmt.Sprintf("%s:%s", env.Env.CacheHost, env.Env.CachePort),
			Password:     env.Env.CachePassword,
			DB:           0,
			ReadTimeout:  time.Duration(env.Env.CacheReadTimeout) * time.Second,
			WriteTimeout: time.Duration(env.Env.CacheWriteTimeout) * time.Second,
		},
	)
	redisOtel.InstrumentTracing(redisClient)
	cmd := redisClient.Ping(ctx)
	if cmd.Err() != nil {
		return nil, fmt.Errorf("ping -> %w", cmd.Err())
	}
	return redisClient, nil
}
