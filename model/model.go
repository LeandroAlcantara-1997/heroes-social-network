package model

import "github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"

type Hero struct {
	Id        string
	HeroName  string
	CivilName string
	Hero      bool
	Universe  string
	Team      *string
}

func New(id string, dto input.HeroRequest) *Hero {
	return &Hero{
		Id:        id,
		HeroName:  dto.HeroName,
		CivilName: dto.CivilName,
		Hero:      dto.Hero,
		Universe:  dto.Universe,
		Team:      dto.Team,
	}
}

type Universe string

const (
	Marvel   Universe = "MARVEL"
	DC       Universe = "DC"
	DCMarvel          = "DC|MARVEL"
)

func CheckUniverse(universe Universe) bool {
	switch universe {
	case Marvel, DC, DCMarvel:
		return true
	}
	return false
}
