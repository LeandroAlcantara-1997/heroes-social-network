package model

import (
	"strings"
	"time"

	hero "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/dto"
)

type Hero struct {
	ID        string
	HeroName  string
	CivilName string
	Hero      bool
	Universe  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	Team      *string
}

func NewHero(id string, dto *hero.HeroRequest) *Hero {
	return &Hero{
		ID:        id,
		HeroName:  strings.ToLower(dto.HeroName),
		CivilName: strings.ToLower(dto.CivilName),
		Hero:      dto.Hero,
		Universe:  dto.Universe,
		CreatedAt: time.Now().UTC(),
		Team:      dto.Team,
	}
}
