package log

import (
	"go-tg-stickfind/internal/config"
	"io"
	"log/slog"
	"os"
)

func NewLogger(buildMode string, isConsoleLogger bool, logFilePath string) (*slog.Logger, io.WriteCloser) {
	var logWriter io.WriteCloser
	if isConsoleLogger {
		logWriter = os.Stdout
	} else {
		var err error
		logWriter, err = os.Create(logFilePath)
		if err != nil {
			panic("Error in opening log file")
		}
	}

	var logger *slog.Logger
	switch buildMode {
	case config.Prod:
		logger = NewProdLogger(logWriter)
	case config.Dev:
		logger = NewDevLogger(logWriter)
	default:
		panic("Error to initialize logger in main")
	}

	return logger, logWriter
}

func NewDevLogger(w io.Writer) *slog.Logger {
	if w == nil {
		panic("error to initialize production logger")
	}

	return slog.New(
		slog.NewTextHandler(w, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
		),
	)
}

func NewProdLogger(w io.Writer) *slog.Logger {
	if w == nil {
		panic("error to initialize production logger")
	}

	return slog.New(
		slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
		),
	)
}
