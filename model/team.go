package model

import (
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
)

type Team struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Universe  string `json:"universe"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewTeam(id string, createdAt time.Time, dto *team.TeamRequest) *Team {
	return &Team{
		ID:        id,
		Name:      dto.Name,
		Universe:  dto.Universe,
		CreatedAt: createdAt,
	}
}
