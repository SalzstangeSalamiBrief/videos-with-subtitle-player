package models

import (
	"log/slog"
	"strings"
)

const (
	LevelTrace = slog.Level(-8)
	LevelFatal = slog.Level(12)
)

var LevelNames = map[slog.Leveler]string{
	LevelTrace: "TRACE",
	LevelFatal: "FATAL",
}

func TransformLevelStringToLeveler(input string) slog.Leveler {
	switch strings.ToLower(input) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "trace":
		return LevelTrace
	case "fatal":
		return LevelFatal
	default:
		return slog.LevelInfo
	}
}
