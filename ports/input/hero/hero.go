package hero

import (
	"context"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
)

type HeroRequest struct {
	HeroName  string  `json:"heroName"`
	CivilName string  `json:"civilName"`
	Hero      bool    `json:"hero"`
	Universe  string  `json:"universe"`
	Team      *string `json:"team,omitempty"`
}

type HeroResponse struct {
	ID        string             `json:"id"`
	HeroName  string             `json:"heroName"`
	CivilName string             `json:"civilName"`
	Hero      bool               `json:"hero"`
	Universe  string             `json:"universe"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at"`
	Team      *team.TeamResponse `json:"team,omitempty"`
}

func NewHeroResponse(id, heroName, civilName, universe string,
	hero bool, createdAt time.Time, updatedAt *time.Time,
	team *team.TeamResponse) *HeroResponse {
	return &HeroResponse{
		ID:        id,
		HeroName:  heroName,
		CivilName: civilName,
		Hero:      hero,
		Universe:  universe,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Team:      team,
	}
}

//go:generate mockgen -destination ../../../mock/hero_mock.go -package=mock -source=hero.go
type Hero interface {
	RegisterHero(ctx context.Context, dto *HeroRequest) (*HeroResponse, error)
	UpdateHero(ctx context.Context, id string, dto *HeroRequest) (*HeroResponse, error)
	GetHeroByID(ctx context.Context, id string) (*HeroResponse, error)
	DeleteHeroByID(ctx context.Context, id string) (err error)
}
