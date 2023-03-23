package input

import "context"

type HeroRequest struct {
	HeroName  string `json:"heroName"`
	CivilName string `json:"civilName"`
	Hero      bool   `json:"hero"`
	Universe  string `json:"universe"`
}

type HeroResponse struct {
	Id        string `json:"id"`
	HeroName  string `json:"heroName"`
	CivilName string `json:"civilName"`
	Hero      bool   `json:"hero"`
	Universe  string `json:"universe"`
}

type Hero interface {
	RegisterHero(ctx context.Context, dto HeroRequest) (*HeroResponse, error)
}
