package log

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/david-fitzgerald/sympal/internal/config"
)

var logger *slog.Logger

func Init() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	logPath := filepath.Join(homeDir, ".sympal", "sympal.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	level := parseLevel(config.Current.LogLevel)
	handler := slog.NewTextHandler(file, &slog.HandlerOptions{Level: level})
	logger = slog.New(handler)

	return nil
}

func parseLevel(s string) slog.Level {
	switch s {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}
