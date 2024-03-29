package response

import (
	"errors"
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
)

var restErrorMap = map[error]int{
	exception.ErrInternalServer:   http.StatusInternalServerError,
	exception.ErrInvalidFields:    http.StatusBadRequest,
	exception.ErrHeroNotFound:     http.StatusNotFound,
	exception.ErrTeamNotFound:     http.StatusNotFound,
	exception.ErrGameNotFound:     http.StatusNotFound,
	exception.ErrInvalidRequest:   http.StatusForbidden,
	exception.ErrTeamAlredyExists: http.StatusBadRequest,
	exception.ErrHeroAlredyExists: http.StatusBadRequest,
	exception.ErrAbilityNotFound:  http.StatusNotFound,
}

func RestError(key error) (int, error) {
	var err *exception.ErrorWithTrace

	if errors.As(key, &err) {
		key = err.GetError()
	}
	code := restErrorMap[key]
	if code != 0 {
		return code, New(key)
	}

	return http.StatusInternalServerError, New(exception.ErrInternalServer)
}

type Exception struct {
	Key string `json:"key"`
}

func New(err error) *Exception {
	return &Exception{
		Key: err.Error(),
	}
}

func (e *Exception) Error() string {
	return e.Key
}
