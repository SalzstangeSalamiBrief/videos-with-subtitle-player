package middleware

import (
	"fmt"
	"net/http"
)

func RequestLoggerHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO CHANGE AFTER ADDING LOGGING
		fmt.Printf("Request URL: %s\n", r.URL.Path)
		fmt.Printf("Request Method: %s\n", r.Method)
		fmt.Printf("Request Origin: %s\n", r.Header.Get("Origin"))
		next.ServeHTTP(w, r)
	})
}
