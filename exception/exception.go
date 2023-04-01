package exception

import (
	"errors"
	"net/http"
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
	ErrInternalServer = errors.New("error.0001")
	ErrInvalidFields  = errors.New("error.0002")
	ErrHeroNotFound   = errors.New("error.0003")
)

var restErrorMap = map[error]int{
	ErrInternalServer: http.StatusInternalServerError,
	ErrInvalidFields:  http.StatusBadRequest,
	ErrHeroNotFound:   http.StatusNotFound,
}

func RestError(key error) (int, error) {
	code := restErrorMap[key]
	if code != 0 {
		return code, New(key.Error())
	}
	return http.StatusInternalServerError, errors.New("internal server error")
}
