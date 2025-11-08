package middlewares

import (
	"backend/internal/problemDetailsErrors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CorsMiddleWareConfiguration struct {
	AllowedCors string
}

type CorsMiddleWareBuilder struct {
	configuration *CorsMiddleWareConfiguration
}

func NewCorsMiddleWare() *CorsMiddleWareBuilder {
	return &CorsMiddleWareBuilder{
		configuration: &CorsMiddleWareConfiguration{},
	}
}

func (builder *CorsMiddleWareBuilder) AddConfiguration(configuration CorsMiddleWareConfiguration) *CorsMiddleWareBuilder {
	builder.configuration = &configuration
	return builder
}

func (builder *CorsMiddleWareBuilder) Build() func(next http.HandlerFunc) http.HandlerFunc {
	builder.configuration.validateConfiguration()

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			allowAllOrigins := builder.configuration.AllowedCors == "*"
			if allowAllOrigins {
				w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			}

			origin := r.Header.Get("Origin")
			fmt.Printf("Origin: %s\n", origin)

			hasMatchedOrigin := false
			allowedCors := strings.Split(builder.configuration.AllowedCors, ",")
			for _, allowedOrigin := range allowedCors {
				if origin == allowedOrigin {
					w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
					hasMatchedOrigin = true
					break
				}
			}

			if origin != "" && !hasMatchedOrigin {
				problemDetailsErrors.NewForbiddenProblemDetails("The origin is not allowed (CORS)").SendErrorResponse(w)
				return
			}

			next(w, r)
		}
	}
}

func (configuration CorsMiddleWareConfiguration) validateConfiguration() {
	if configuration.AllowedCors == "" {
		log.Fatal("CorsMiddleware requires the field allowedCors to be set")
	}
}
