package problemDetailsErrors

import (
	"encoding/json"
	"log"
	"net/http"
)

// ProblemDetailsError The type is based on RFC-7807
//
// https://datatracker.ietf.org/doc/html/rfc7807
type ProblemDetailsError struct {
	// ProblemDetailsError.Title The string representation of the HTTP-Status code
	Title string `json:"title"`
	// ProblemDetailsError.Status The HTTP-Status code
	Status int `json:"status"`
	// ProblemDetailsError.Detail A human-readable explanation of the error
	Detail string `json:"detail"`
	// ProblemDetailsError.ProblemType This field references corresponding sections in the RFC-9110
	// https://datatracker.ietf.org/doc/html/rfc9110
	ProblemType string `json:"type"`
}

func (e *ProblemDetailsError) Error() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("Error marshalling ProblemDetailsError: %v\n", err)
	}

	return string(bytes)
}

func (e *ProblemDetailsError) SendErrorResponse(w http.ResponseWriter) {
	// TODO PROPPER LOGGING
	log.Printf("SendErrorResponse %v\n", e)
	http.Error(w, e.Error(), e.Status)
}
