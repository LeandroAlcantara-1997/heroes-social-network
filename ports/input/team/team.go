package team

import (
	"context"
	"time"
)

type TeamRequest struct {
	Name     string `json:"name"`
	Universe string `json:"universe"`
}

type TeamResponse struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Universe  string     `json:"universe"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func NewTeamResponse(id, name, universe string,
	createdAt time.Time, updatedAt *time.Time) *TeamResponse {
	return &TeamResponse{
		Id:        id,
		Name:      name,
		Universe:  universe,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

//go:generate mockgen -destination ../../../mock/team_mock.go -package=mock -source=team.go
type Team interface {
	RegisterTeam(ctx context.Context, team *TeamRequest) (*TeamResponse, error)
}
