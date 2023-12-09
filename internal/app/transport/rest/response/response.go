package response

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
)

var restErrorMap = map[error]int{
	exception.ErrInternalServer:   http.StatusInternalServerError,
	exception.ErrInvalidFields:    http.StatusBadRequest,
	exception.ErrHeroNotFound:     http.StatusNotFound,
	exception.ErrTeamNotFound:     http.StatusNotFound,
	exception.ErrInvalidRequest:   http.StatusForbidden,
	exception.ErrTeamAlredyExists: http.StatusBadRequest,
	exception.ErrHeroAlredyExists: http.StatusBadRequest,
}

func RestError(key error) (int, error) {
	code := restErrorMap[key]
	if code != 0 {
		return code, exception.New(key.Error())
	}
	return http.StatusInternalServerError, exception.ErrInternalServer
}
