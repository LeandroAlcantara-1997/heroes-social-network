package exception

import (
	"errors"
)

type Exception struct {
	Key string `json:"key"`
}

func New(key string) *Exception {
	return &Exception{
		Key: key,
	}
}

func (e *Exception) Error() string {
	return e.Key
}

var (
	ErrInternalServer   = errors.New("error.0001")
	ErrInvalidFields    = errors.New("error.0002")
	ErrHeroNotFound     = errors.New("error.0003")
	ErrTeamNotFound     = errors.New("error.0004")
	ErrInvalidRequest   = errors.New("error.0005")
	ErrTeamAlredyExists = errors.New("error.0006")
	ErrHeroAlredyExists = errors.New("error.0007")
)
