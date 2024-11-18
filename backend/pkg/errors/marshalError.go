package errors

import (
	"fmt"
	"net/http"
)

type MarshalError struct {
	InnerError string
}

func (e *MarshalError) Error() string {
	return fmt.Sprintf("Could not marshal file tree - inner error '%v'", e.InnerError)
}

func (e *MarshalError) StatusCode() int {
	return http.StatusInternalServerError
}
