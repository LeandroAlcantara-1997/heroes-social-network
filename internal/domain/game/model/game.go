package model

import (
	"time"

	console "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/pkg/universe"
)

type Game struct {
	ID          string
	Name        string
	TeamID      *string
	HeroID      *string
	ReleaseYear int
	Universe    universe.Universe
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	Consoles    []console.Console
}

func NewGame(id string, req *dto.GameRequest) *Game {
	return &Game{
		ID:          id,
		Universe:    req.Universe,
		Name:        req.Name,
		ReleaseYear: req.ReleaseYear,
		HeroID:      req.HeroID,
		TeamID:      req.TeamID,
		CreatedAt:   time.Now().UTC(),
		Consoles:    req.Consoles,
	}
}
