package model

import (
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
}

func NewGame(id string, req *dto.GameRequest) *Game {
	return &Game{
		ID:       id,
		Universe: req.Universe,
		Name:     req.Name,
	}
}
