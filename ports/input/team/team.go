package team

import (
	"context"
	"time"
)

type TeamRequest struct {
	Name     string `json:"name" validate:"required"`
	Universe string `json:"universe" validate:"universe,required"`
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

//go:generate mockgen -destination ../../../mock/team_mock.go -package=mock -source=team.go
type Team interface {
	RegisterTeam(ctx context.Context, dto *TeamRequest) (*TeamResponse, error)
	GetTeamByID(ctx context.Context, id string) (*TeamResponse, error)
	DeleteTeamByID(ctx context.Context, id string) error
}
