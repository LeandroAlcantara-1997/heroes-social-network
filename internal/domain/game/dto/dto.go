package dto

import (
	"errors"
	"time"

	console "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/pkg/universe"
)

type GameRequest struct {
	Name        string            `json:"name" example:"X-Men Legends"`
	ReleaseYear int               `json:"releaseYear" example:"2004"`
	TeamID      *string           `json:"teamId" example:"c184abee-d573-442d-b1b7-ba93aff61fb6"`
	HeroID      *string           `json:"heroId" example:"c184abee-d573-442d-b1b7-ba93aff61fb6"`
	Universe    universe.Universe `json:"universe" example:"MARVEL"`
	Consoles    []console.Console `json:"consoles" example:"Playstation2"`
}

type GameResponse struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	ReleaseYear int               `json:"releaseYear"`
	Universe    universe.Universe `json:"universe"`
	TeamID      *string           `json:"teamId,omitempty"`
	HeroID      *string           `json:"heroId,omitempty"`
	CreatedAt   time.Time         `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time        `json:"updatedAt,omitempty"`
	Consoles    []console.Console `json:"consoles"`
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

	for c := range g.Consoles {
		if g.Consoles[c] == "" {
			return errors.New("invalid console")
		}
	}
	return nil
}

func NewGameResponse(id, name string, releaseYear int, teamID, heroID *string, universe universe.Universe,
	createdAt time.Time, updatedAt *time.Time, consoles []console.Console) *GameResponse {
	return &GameResponse{
		ID:          id,
		Name:        name,
		ReleaseYear: releaseYear,
		TeamID:      teamID,
		HeroID:      heroID,
		Universe:    universe,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Consoles:    consoles,
	}
}
