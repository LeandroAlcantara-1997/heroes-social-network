package model

import (
	"strings"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/dto"
)

type Team struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Universe  string `json:"universe"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewTeam(id string, createdAt time.Time, dto *dto.TeamRequest) *Team {
	return &Team{
		ID:        id,
		Name:      strings.ToLower(dto.Name),
		Universe:  dto.Universe,
		CreatedAt: createdAt,
	}
}

type Universe string

const (
	Marvel   Universe = "MARVEL"
	DC       Universe = "DC"
	DCMarvel Universe = "DC|MARVEL"
)

func CheckUniverse(universe Universe) bool {
	universe = Universe(strings.ToUpper(string(universe)))
	switch universe {
	case Marvel, DC, DCMarvel:
		return true
	}
	return false
}
