package main

import (
	"backend/cmd/logger/models"
	"backend/cmd/logger/utilities"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"
	"runtime/debug"
	"strings"
)

type CustomHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type CustomHandler struct {
	slog.Handler
	l *log.Logger
}

var logger *slog.Logger

// TODO
// TODO ATTRIBUTES: ADD SOURCE, JSON, MIN LEVEL?
func main( /* TODO MIN LEVEL AS PARAM */ ) {
	//logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: LevelTrace, ReplaceAttr: replaceAttr}))
	minLevel := models.LevelTrace // slog.LevelWarn
	shouldAddSource := true       // TODO PARAM
	options := CustomHandlerOptions{SlogOpts: slog.HandlerOptions{AddSource: shouldAddSource, Level: minLevel, ReplaceAttr: replaceAttr}}
	handler := NewCustomHandler(os.Stdout, options)
	logger = slog.New(handler)

	//slog.SetDefault(logger)
	//logger.Info("Hello world", "user", 134, "baka", []string{"a", "b"})
	//slog.Info("IAM THJE DEFAULT BAKA")
	buildInfo, _ := debug.ReadBuildInfo()
	logger.Warn("Warn", slog.String("go_version", buildInfo.GoVersion))
	logger.Error("Error")
	logger.Info("Info")
	logger.Debug("Debug")
	//childLogger := logger.With(slog.Group("Program_properties", slog.Int("pid", os.Getpid()), slog.String("go_version", buildInfo.GoVersion)))
	//logger.LogAttrs(nil, slog.LevelInfo, "hello, world",
	//	slog.Group("Program properties", slog.Int("pid", os.Getpid()), slog.String("go_version", buildInfo.GoVersion)))
	//childLogger.LogAttrs(nil, slog.LevelError, "Hello from the child", slog.Int("Attrs", 1))
	// src: https://go.dev/blog/slog
	// TODO USE BUILDER PATTERN
	// TODO INJECT INTO ROUTER OR EXPORT REFERENCE?
}

func replaceAttr(groups []string, attributes slog.Attr) slog.Attr {
	if attributes.Key == slog.LevelKey {
		level := attributes.Value.Any().(slog.Level)
		label, exists := models.LevelNames[level]
		if !exists {
			label = level.String()
		}

		attributes.Value = slog.StringValue(label)
	}

	return attributes
}

// TODO CUSTOM HANDLER

func NewCustomHandler(
	out io.Writer,
	opts CustomHandlerOptions,
) *CustomHandler {
	h := &CustomHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}

func (h *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	level := models.TransformLevelStringToLeveler(r.Level.String()).Level()
	marshalString := transformMessageToJsonString(*h, r)
	colorizedMessage := utilities.ColorizeMessageByLevel(level, level.String(), r.Message, marshalString)
	h.l.Println(colorizedMessage...)
	return nil
}

func transformMessageToJsonString(h CustomHandler, r slog.Record) string {
	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	marshalResult, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		h.l.Fatalln("Could not create a marshal JSON", err.Error())
	}

	compactJsonDestination := bytes.Buffer{}
	err = json.Compact(&compactJsonDestination, marshalResult)
	if err != nil {
		h.l.Fatalln("Could not create a compact JSON string", err.Error())
	}

	return compactJsonDestination.String()
}

// TODO UTILITIES
// TODO https://betterstack.com/community/guides/logging/logging-in-go/
func getHandler(handlerType string, minLevel string) slog.Handler {
	options := slog.HandlerOptions{AddSource: true, Level: models.TransformLevelStringToLeveler(minLevel)}
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
