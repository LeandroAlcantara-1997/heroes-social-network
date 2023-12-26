package dto

import (
	"errors"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/pkg/universe"
)

type GameRequest struct {
	Name        string            `json:"name"`
	ReleaseYear int               `json:"releaseYear"`
	TeamID      *string           `json:"teamId"`
	HeroID      *string           `json:"heroId"`
	Universe    universe.Universe `json:"universe"`
}

type GameResponse struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	ReleaseYear int               `json:"releaseYear"`
	Universe    universe.Universe `json:"universe"`
	TeamID      *string           `json:"teamId,omitempty"`
	HeroID      *string           `json:"heroId,omitempty"`
}

func (g *GameRequest) Validator() error {
	if g.Name == "" {
		return errors.New("invalid name")
	}

	if g.ReleaseYear < 1975 {
		return errors.New("invalid release year")
	}

	if (g.TeamID == nil && g.HeroID == nil) || (g.TeamID != nil && g.HeroID != nil) {
		return errors.New("invalid teamId or heroId")
	}

	if !universe.CheckUniverse(g.Universe) {
		return errors.New("invalid universe")
	}

	return nil
}
