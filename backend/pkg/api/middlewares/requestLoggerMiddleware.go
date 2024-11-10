package middlewares

import (
	"fmt"
	"net/http"
)

func RequestLoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO CHANGE AFTER ADDING LOGGING
		fmt.Printf("Request URL: %s\n", r.URL.Path)
		fmt.Printf("Request Method: %s\n", r.Method)
		fmt.Printf("Request Origin: %s\n", r.Header.Get("Origin"))
		next(w, r)
	}
}
