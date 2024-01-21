package exception

import (
	"errors"
)

var (
	ErrInternalServer   = errors.New("error.0001")
	ErrInvalidFields    = errors.New("error.0002")
	ErrHeroNotFound     = errors.New("error.0003")
	ErrTeamNotFound     = errors.New("error.0004")
	ErrInvalidRequest   = errors.New("error.0005")
	ErrTeamAlredyExists = errors.New("error.0006")
	ErrHeroAlredyExists = errors.New("error.0007")
	ErrGameNotFound     = errors.New("error.0008")
)

type ErrorWithTrace struct {
	trace string
	err   error
}

func (e *ErrorWithTrace) GetError() error {
	return e.err
}

func (e *ErrorWithTrace) Error() string {
	return e.err.Error()
}
func New(trace string, err error) *ErrorWithTrace {
	var errTrace *ErrorWithTrace
	if errors.As(err, &errTrace) {
		return &ErrorWithTrace{
			trace: errTrace.trace + trace,
			err:   err,
		}
	}

	return &ErrorWithTrace{
		trace: trace,
		err:   err,
	}
}
