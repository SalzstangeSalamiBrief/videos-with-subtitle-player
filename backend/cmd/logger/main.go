package main

import (
	"context"
	"github.com/fatih/color"
	"io"
	"log"
	"log/slog"
	"os"
	"runtime/debug"
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

type ColorHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type ColorHandler struct {
	slog.Handler
	l *log.Logger
}

var logger *slog.Logger

// TODO
// TODO ATTRIBUTES: ADD SOURCE, JSON, MIN LEVEL?
func main() {
	//logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: LevelTrace, ReplaceAttr: replaceAttr}))
	logger = slog.New(NewCustomColorHandler(os.Stdout, ColorHandlerOptions{SlogOpts: slog.HandlerOptions{AddSource: true, Level: LevelTrace, ReplaceAttr: replaceAttr}}))

	//slog.SetDefault(logger)
	//logger.Info("Hello world", "user", 134, "baka", []string{"a", "b"})
	//slog.Info("IAM THJE DEFAULT BAKA")
	buildInfo, _ := debug.ReadBuildInfo()
	childLogger := logger.With(slog.Group("Program_properties", slog.Int("pid", os.Getpid()), slog.String("go_version", buildInfo.GoVersion)))
	//logger.LogAttrs(nil, slog.LevelInfo, "hello, world",
	//	slog.Group("Program properties", slog.Int("pid", os.Getpid()), slog.String("go_version", buildInfo.GoVersion)))
	childLogger.LogAttrs(nil, slog.LevelError, "Hello from the child", slog.Int("Attrs", 1))
	// src: https://go.dev/blog/slog
	// TODO USE BUILDER PATTERN
	// TODO INJECT INTO ROUTER OR EXPORT REFERENCE?
}

func replaceAttr(groups []string, attributes slog.Attr) slog.Attr {
	if attributes.Key == slog.LevelKey {
		level := attributes.Value.Any().(slog.Level)
		label, exists := LevelNames[level]
		if !exists {
			label = level.String()
		}

		attributes.Value = slog.StringValue(label)
	}

	return attributes
}

// TODO https://betterstack.com/community/guides/logging/logging-in-go/
func getHandler(handlerType string, minLevel string) slog.Handler {
	options := slog.HandlerOptions{AddSource: true, Level: getOutputLevel(minLevel)}
	if handlerType == "" && os.Getenv("env") == strings.ToLower("prod") {
		handlerType = "json"
	}

	switch strings.ToLower(handlerType) {
	case "json":
		return slog.NewJSONHandler(os.Stdout, &options)
	default:
		return slog.NewTextHandler(os.Stdout, &options)
	}
}

func getOutputLevel(input string) slog.Leveler {
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

func NewCustomColorHandler(
	out io.Writer,
	opts ColorHandlerOptions,
) *ColorHandler {
	h := &ColorHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}

func (h *ColorHandler) Handle(ctx context.Context, record slog.Record) error {
	logLevel := record.Level
	message := record.Message
	switch logLevel {
	case slog.LevelWarn:
		message = color.YellowString(message)
		break
	case slog.LevelError:
		message = color.RedString(message)
		break
	case LevelTrace:
		message = color.BlueString(message)
		break
	case LevelFatal:
		message = color.MagentaString(message)
		break
	default:
		message = color.WhiteString(message)
		break
	}

	return nil
}
