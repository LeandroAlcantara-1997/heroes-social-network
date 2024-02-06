package dto

import (
	"errors"
	"time"
)

type AbilityRequest struct {
	Description string `json:"description"`
}

type AbilityResponse struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

func (a *AbilityRequest) Validator() error {
	if a.Description == "" {
		return errors.New("invalid description")
	}

	return nil
}

func NewAbilityResponse(id, description string, createdAt time.Time, updatedAt *time.Time) *AbilityResponse {
	return &AbilityResponse{
		ID:          id,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
