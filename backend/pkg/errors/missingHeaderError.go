package errors

import (
	"fmt"
	"net/http"
)

type MissingHeaderError struct {
	Id     string
	Header string
}

func (e *MissingHeaderError) Error() string {
	return fmt.Sprintf("Missing Header: Could not get file with id '%v' because the header '%v' is missing", e.Id, e.Header)
}

func (e *MissingHeaderError) StatusCode() int {
	return http.StatusBadRequest
}
