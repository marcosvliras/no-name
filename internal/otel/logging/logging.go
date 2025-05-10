package logging

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func GetLogger() *slog.Logger {
	env := os.Getenv("ENVIRONMENT")
	if env == "PROD" {
		Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
		return Logger
	}
	Logger = slog.Default()
	return Logger
}
