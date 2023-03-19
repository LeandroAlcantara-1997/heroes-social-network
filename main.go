package main

import (
	"log"

	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	migrations()
}

func migrations() {
	m, err := migrate.New(
		"migration",
		config.Env.DataBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
