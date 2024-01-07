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

type errorWithTrace struct {
	trace string
	err   error
}

func (e *errorWithTrace) GetError() error {
	return e.err
}

func (e *errorWithTrace) Error() string {
	return e.err.Error()
}
func New(trace string, err error) *errorWithTrace {
	var errTrace *errorWithTrace
	if errors.As(err, &errTrace) {
		return &errorWithTrace{
			trace: errTrace.trace + trace,
			err:   err,
		}
	}

	return &errorWithTrace{
		trace: trace,
		err:   err,
	}
}
