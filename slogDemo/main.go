package main

import (
	"log/slog"
	"os"
)

func main() {
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
