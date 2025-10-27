package problemDetailsErrors

import (
	"net/http"
)

func NewNotFoundProblemDetails(detail string) *ProblemDetailsError {
	return &ProblemDetailsError{
		Status:      http.StatusNotFound,
		Title:       "Not found",
		Detail:      detail,
		ProblemType: "https://datatracker.ietf.org/doc/html/rfc9110#section-15.5.5",
	}
}

func NewBadRequestProblemDetails(detail string) *ProblemDetailsError {
	return &ProblemDetailsError{
		Status:      http.StatusBadRequest,
		Title:       "Bad request",
		Detail:      detail,
		ProblemType: "https://datatracker.ietf.org/doc/html/rfc9110#section-15.5.1",
	}
}

func NewInternalServerErrorProblemDetails(detail string) *ProblemDetailsError {
	return &ProblemDetailsError{
		Status:      http.StatusInternalServerError,
		Title:       "Internal server error",
		Detail:      detail,
		ProblemType: "https://datatracker.ietf.org/doc/html/rfc9110#section-15.6.1",
	}
}

func NewForbiddenProblemDetails(detail string) *ProblemDetailsError {
	return &ProblemDetailsError{
		Status:      http.StatusForbidden,
		Title:       "Forbidden",
		Detail:      detail,
		ProblemType: "https://datatracker.ietf.org/doc/html/rfc9110#section-15.5.4",
	}
}
