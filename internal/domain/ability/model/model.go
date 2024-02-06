package model

import (
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/dto"
	"github.com/google/uuid"
)

type Ability struct {
	ID          string
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func NewAbility(req *dto.AbilityRequest) *Ability {
	return &Ability{
		ID:          uuid.NewString(),
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
}
