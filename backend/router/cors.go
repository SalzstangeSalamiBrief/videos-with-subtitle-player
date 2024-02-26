package router

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func useCors(w http.ResponseWriter, r *http.Request) {
	configuredCors := os.Getenv("ALLOWED_CORS")
	if configuredCors == "" {
		return
	}

	allowAllOrigins := configuredCors == "*"
	if allowAllOrigins {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	}

	allowedCors := strings.Split(configuredCors, ",")
	for _, allowedOrigin := range allowedCors {
		origin := r.Header.Get("Origin")
		fmt.Printf("Origin: %s\n", origin)
		if origin == allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			break
		}
	}
}
