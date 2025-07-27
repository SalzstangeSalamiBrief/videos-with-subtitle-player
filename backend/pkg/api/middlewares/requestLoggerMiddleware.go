package middlewares

import (
	"log"
	"net/http"
	"os"
)

type RequestLoggerBuilder struct {
	logger *log.Logger
}

func NewRequestLogger() *RequestLoggerBuilder {
	return &RequestLoggerBuilder{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (builder *RequestLoggerBuilder) SetLogger(logger *log.Logger) *RequestLoggerBuilder {
	builder.logger = logger
	return builder
}

func (builder *RequestLoggerBuilder) Build() func(next http.HandlerFunc) http.HandlerFunc {
	logger := builder.logger

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// TODO CHANGE AFTER ADDING LOGGING
			logger.Printf("Request URL: %s\n", r.URL.Path)
			logger.Printf("Request Method: %s\n", r.Method)
			logger.Printf("Request Origin: %s\n", r.Header.Get("Origin"))
			next(w, r)
		}
	}
}
