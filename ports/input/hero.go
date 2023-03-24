package input

import (
	"context"
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
	Team      *TeamResponse `json:"team,omitempty"`
}

type Hero interface {
	RegisterHero(ctx context.Context, dto *HeroRequest) (*HeroResponse, error)
}
