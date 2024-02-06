package dto

import (
	"errors"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/pkg/universe"
)

type HeroRequest struct {
	HeroName  string  `json:"heroName" validate:"required"`
	CivilName string  `json:"civilName" validate:"required"`
	Hero      bool    `json:"hero" validate:"required"`
	Universe  string  `json:"universe" validate:"universe,required"`
	Team      *string `json:"team,omitempty"`
}

type HeroResponse struct {
	ID        string     `json:"id"`
	HeroName  string     `json:"heroName"`
	CivilName string     `json:"civilName"`
	Hero      bool       `json:"hero"`
	Universe  string     `json:"universe"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Team      *string    `json:"team,omitempty"`
}

func NewHeroResponse(id, heroName, civilName, universe string,
	hero bool, createdAt time.Time, updatedAt *time.Time,
	team *string) *HeroResponse {
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

func (h *HeroRequest) Validator() error {
	if h.CivilName == "" {
		return errors.New("invalid civil name")
	}

	if h.HeroName == "" {
		return errors.New("invalid hero name")
	}

	if h.Universe == "" || !universe.CheckUniverse(universe.Universe(h.Universe)) {
		return errors.New("invalid universe")
	}

	return nil
}

type AddAbilityToHeroRequest struct {
	AbilityID string `form:"ability" binding:"required,uuid"`
	HeroID    string `form:"hero" binding:"required,uuid"`
}
