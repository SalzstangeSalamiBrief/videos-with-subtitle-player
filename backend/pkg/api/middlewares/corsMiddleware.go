package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func CorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		next.ServeHTTP(w, r)
	})
}

// TODO DO SOMETHING LIKE THIS TO ADD MULTIPLE MIDDLEWARES
//func myHanlder(middlewares ...http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//
//		for _, m := range middlewares {
//			next = m(next)
//		}
//
//		return next
//	}
//
//}
