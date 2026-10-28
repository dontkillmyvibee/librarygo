package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func New(cfg Config) (*Logger, error) {
	var level slog.Level

	if err := level.UnmarshalText([]byte(cfg.LogLevel)); err != nil {
		return nil, fmt.Errorf("parse log level %q: %w", cfg.LogLevel, err)
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return &Logger{
		logger: slog.New(handler),
	}, nil
}

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		logger: l.logger.With(args...),
	}
}
