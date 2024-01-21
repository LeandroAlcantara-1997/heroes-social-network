package dto

import (
	"errors"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/pkg/universe"
)

type TeamRequest struct {
	Name     string `json:"name" validate:"required"`
	Universe string `json:"universe" validate:"universe,required"`
}

type GetTeamByName struct {
	Name string `uri:"name" binding:"required"`
}

type TeamResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Universe  string     `json:"universe"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func NewTeamResponse(id, name, universe string,
	createdAt time.Time, updatedAt *time.Time) *TeamResponse {
	return &TeamResponse{
		ID:        id,
		Name:      name,
		Universe:  universe,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (t *TeamRequest) Validator() error {
	if t.Name == "" {
		return errors.New("invalid team name")
	}

	if !universe.CheckUniverse(universe.Universe(t.Universe)) {
		return errors.New("invalid universe")
	}

	return nil
}
