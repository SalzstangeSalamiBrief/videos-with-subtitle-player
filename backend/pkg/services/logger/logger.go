package logger

import "log/slog"

var logger *slog.Logger

// TODO
func LoggerNew() {
	logger = slog.Default()
	// src: https://go.dev/blog/slog
	// TODO USE BUILDER PATTERN
	// TODO INJECT INTO ROUTER OR EXPORT REFERENCE?
}
