package errors

import (
	"fmt"
	"net/http"
)

type FileNotFoundError struct {
	Id string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("File not found: Could not get file with id '%v'", e.Id)
}

func (e *FileNotFoundError) StatusCode() int {
	return http.StatusBadRequest
}
