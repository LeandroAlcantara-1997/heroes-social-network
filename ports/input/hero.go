package input

import (
	"context"
	"time"
)

type HeroRequest struct {
	HeroName  string  `json:"heroName"`
	CivilName string  `json:"civilName"`
	Hero      bool    `json:"hero"`
	Universe  string  `json:"universe"`
	Team      *string `json:"team"`
}

type HeroResponse struct {
	Id        string        `json:"id"`
	HeroName  string        `json:"heroName"`
	CivilName string        `json:"civilName"`
	Hero      bool          `json:"hero"`
	Universe  string        `json:"universe"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at"`
	Team      *TeamResponse `json:"team,omitempty"`
}

func NewHeroResponse(id, heroName, civilName, universe string,
	hero bool, created_at time.Time, updated_at *time.Time,
	team *TeamResponse) *HeroResponse {
	return &HeroResponse{
		Id:        id,
		HeroName:  heroName,
		CivilName: civilName,
		Hero:      hero,
		Universe:  universe,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		Team:      team,
	}
}

//go:generate mockgen -destination ../../mock/hero_mock.go -package=mock -source=hero.go
type Hero interface {
	RegisterHero(ctx context.Context, dto *HeroRequest) (*HeroResponse, error)
	UpdateHero(ctx context.Context, id string, dto *HeroRequest) (*HeroResponse, error)
	GetHeroByID(ctx context.Context, id string) (*HeroResponse, error)
	DeleteHeroByID(ctx context.Context, id string) (err error)
}
