package dto

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameRequestValidatorSuccess(t *testing.T) {
	g := &GameRequest{
		Name:        "Batman: Arkham",
		ReleaseYear: 2023,
		TeamID:      nil,
		HeroID:      []string{"123456"},
		Universe:    "DC",
	}
	err := g.Validator()
	assert.NoError(t, err)
}

func TestGameRequestValidatorFailInvalidName(t *testing.T) {
	var expected = errors.New("invalid name")
	g := &GameRequest{
		Name:        "",
		ReleaseYear: 2023,
		TeamID:      nil,
		HeroID:      []string{"123456"},
		Universe:    "DC",
	}
	err := g.Validator()
	assert.Equal(t, expected.Error(), err.Error())
}

func TestGameRequestValidatorFailInvalidReleaseYear(t *testing.T) {
	var expected = errors.New("invalid release year")
	g := &GameRequest{
		Name:        "Batman: Arkham",
		ReleaseYear: 1000,
		TeamID:      nil,
		HeroID:      []string{"123456"},
		Universe:    "DC",
	}
	err := g.Validator()
	assert.Equal(t, expected.Error(), err.Error())
}
