package internal

import (
	"log/slog"
	"os"
)

// MakeLogger creates a new logger with the given options.
// The logger writes the messages in the format of JSON object to os.Stderr.
func MakeLogger(opts *slog.HandlerOptions) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stderr, opts))

}
