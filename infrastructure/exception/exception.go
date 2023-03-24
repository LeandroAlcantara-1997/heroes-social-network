package exception

import (
	"fmt"
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

const (
	InternalServerError = "error.0001"
	InvalidFieldsError  = "error.0002"
)

var restErrorMap = map[string]int{
	InternalServerError: http.StatusInternalServerError,
	InvalidFieldsError:  http.StatusBadRequest,
}

func RestError(key string) (int, error) {
	code := restErrorMap[key]
	if code != 0 {
		return code, New(key)
	}
	return http.StatusInternalServerError, fmt.Errorf("")
}
