package handlers

import (
	"fmt"
	"net/http"
)

type ErrorHandlerInput interface {
	Error() string
	StatusCode() int
}

func ErrorHandler(w http.ResponseWriter, err ErrorHandlerInput) {
	fmt.Println(err.Error())
	http.Error(w, err.Error(), err.StatusCode())
}
