package logging

import (
	"context"
	"log/slog"
)

type CliLogger struct {
	logger *slog.Logger
}

func NewCliLogger() *CliLogger {
	logger := slog.Default()
	return &CliLogger{
		logger: logger,
	}
}

func (l *CliLogger) Debug(ctx context.Context, msg string) {
	l.logger.Debug(msg)
}
func (l *CliLogger) Info(ctx context.Context, msg string) {
	l.logger.Info(msg)
}
func (l *CliLogger) Warning(ctx context.Context, msg string) {
	l.logger.Warn(msg)
}
func (l *CliLogger) Error(ctx context.Context, msg string) {
	l.logger.Error(msg)
}
func (l *CliLogger) Fatal(ctx context.Context, msg string) {
	l.logger.Error(msg)
}
func (l *CliLogger) Close(ctx context.Context) error {
	return nil
}
