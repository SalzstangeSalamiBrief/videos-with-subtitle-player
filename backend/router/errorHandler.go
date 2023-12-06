package router

import "net/http"

func ErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}
