package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	client *pgxpool.Pool
}

func New(client *pgxpool.Pool) *repository {

	return &repository{
		client: client,
	}
}
