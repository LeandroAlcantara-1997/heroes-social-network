package dto

import (
	"errors"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/model"
)

type ConsoleRequest struct {
	Names []model.Console `json:"consoles" example:"Playstation5"`
}

type ConsoleResponse struct {
	Names []model.Console `json:"consoles"`
}

func (c *ConsoleRequest) Validator() error {
	for n := range c.Names {
		if c.Names[n] == "" {
			return errors.New("invalid console")
		}
	}

	return nil
}
