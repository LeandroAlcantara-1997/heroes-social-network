package model

import "github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"

type Hero struct {
	Id        string
	HeroName  string
	CivilName string
	Hero      bool
	Universe  string
}

func New(dto input.HeroRequest) *Hero {
	return &Hero{
		HeroName:  dto.HeroName,
		CivilName: dto.CivilName,
		Hero:      dto.Hero,
		Universe:  dto.Universe,
	}
}
