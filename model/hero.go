package model

import (
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/hero"
)

type Hero struct {
	Id        string
	HeroName  string
	CivilName string
	Hero      bool
	Universe  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	Team      *string
}

func New(id string, dto *hero.HeroRequest) *Hero {
	return &Hero{
		Id:        id,
		HeroName:  dto.HeroName,
		CivilName: dto.CivilName,
		Hero:      dto.Hero,
		Universe:  dto.Universe,
		CreatedAt: time.Now().UTC(),
		Team:      dto.Team,
	}
}
