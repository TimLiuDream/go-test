package main

import (
	"fmt"
	"log/slog"
	"os"
)

func log() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	jsonLogger := slog.New(jsonHandler).WithGroup("request")
	jsonLogger.Error("",
		slog.String("url", "https://example.com"),
		slog.String("method", "GET"),
		slog.Int("response-code", 200),
	)
}

func Infof(logger *slog.Logger, format string, args ...any) {
	logger.Info(fmt.Sprintf(format, args...))
}
