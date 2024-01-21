package validator

import (
	"github.com/google/uuid"
)

func UUIDValidator(id string) bool {
	if id != "" {
		_, err := uuid.Parse(id)
		return err == nil
	}

	return false
}
