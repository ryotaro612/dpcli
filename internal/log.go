package internal

import (
	"log/slog"
	"os"
)

// NewLogger creates a new logger with the given options.
// The logger writes the messages in the format of JSON object to os.Stderr.
func NewLogger(verbose bool) *slog.Logger {
	var logOptions *slog.HandlerOptions
	if verbose {
		logOptions = &slog.HandlerOptions{Level: slog.LevelDebug}
	}
	return slog.New(slog.NewJSONHandler(os.Stderr, logOptions))
}
