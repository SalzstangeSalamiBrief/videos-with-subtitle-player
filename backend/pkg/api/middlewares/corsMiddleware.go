package middlewares

import (
	"backend/internal/config"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configuredCors := os.Getenv("ALLOWED_CORS")
		if config.AppConfiguration.AllowedCors == "" {
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

		next(w, r)
	}
}
