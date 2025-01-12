package utilities

import (
	"backend/cmd/logger/models"
	"log/slog"
)

type Color string

const (
	DEFAULT Color = "\033[0m"
	YELLOW  Color = "\033[33m"
	RED     Color = "\033[31m"
	GRAY    Color = "\033[90m"
	ORANGE  Color = "\033[38;5;202m"
	WHITE   Color = "\033[37m"
)

func ColorizeMessageByLevel(logLevel slog.Level, content ...any) []any {
	color := getColor(logLevel)
	resetColor := getResetColor()

	result := []any{color.ToString()}
	result = append(result, content...)
	result = append(result, resetColor.ToString())
	return result
}

func (c *Color) ToString() string {
	return string(*c)
}

func getResetColor() Color {
	return DEFAULT
}

func getColor(logLevel slog.Level) Color {
	switch logLevel {
	case slog.LevelWarn:
		return ORANGE
	case slog.LevelError:
		return RED
	case models.LevelTrace:
		return GRAY
	case models.LevelFatal:
		return RED
	case slog.LevelDebug:
		return WHITE
	default:
		return YELLOW
	}
}
