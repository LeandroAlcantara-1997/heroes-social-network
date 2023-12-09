package repository

import (
	"github.com/jackc/pgx/v5"
)

type repository struct {
	client *pgx.Conn
}

func New(client *pgx.Conn) *repository {
	return &repository{
		client: client,
	}
}
