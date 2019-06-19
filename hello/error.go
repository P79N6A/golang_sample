package hello

import (
	"fmt"
	"time"
)

// HelloError is a error implement
type HelloError struct {
	when time.Time
	what string
}

func (e *HelloError) Error() string {
	return fmt.Sprintf("when: %v, what: %v", e.when, e.what)
}

func Run() error {
	return &HelloError{time.Now(), "it's a HelloError"}
}
