package main

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

func main() {
	w := os.Stderr

	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.DateTime,
		}),
	))
	slog.Info("aaaaaaaaa")
}
