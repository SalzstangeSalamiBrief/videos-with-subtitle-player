package errors

import (
	"fmt"
	"net/http"
)

type OsError struct {
	Id         string
	InnerError string
}

func (e *OsError) Error() string {
	return fmt.Sprintf("Could not access file with id '%v' -\n Inner error: '%v'", e.Id, e.InnerError)
}

func (e *OsError) StatusCode() int {
	return http.StatusInternalServerError
}
