package exception

import "time"

type Exception struct {
	Timestamp time.Time `json:"timestamp"`
	Reason    string    `json:"reason"`
}

func (e *Exception) Error() string {
	return ""
}
