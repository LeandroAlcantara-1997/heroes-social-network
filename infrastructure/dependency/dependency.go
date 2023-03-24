package dependency

import (
	"context"
	"fmt"

	"github.com/LeandroAlcantara-1997/heroes-social-network/domain/heroes"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/config"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/repository"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	"github.com/jackc/pgx/v5"
)

type dependency struct {
	HeroUseCase input.Hero
}

type components struct {
	clientPgx *pgx.Conn
}

func LoadDependency(ctx context.Context) (*dependency, error) {
	cmp, err := loadExternalTools(ctx)
	if err != nil {
		return nil, err
	}
	heroService := heroes.New(
		repository.New(cmp.clientPgx),
	)

	return &dependency{
		HeroUseCase: heroService,
	}, nil
}

func loadExternalTools(ctx context.Context) (*components, error) {
	clientPgx, err := pgx.Connect(
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

	return &components{
		clientPgx: clientPgx,
	}, nil
}
