package logging

//var Logger *slog.Logger
//
//func GetLogger() *slog.Logger {
//	env := os.Getenv("ENVIRONMENT")
//	if env == "PROD" {
//		Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
//		return Logger
//	}
//	Logger = slog.Default()
//	return Logger
//}

import (
	"context"
	"time"

	"github.com/marcosvliras/sophie/internal/otel/config"

	"go.opentelemetry.io/otel/log"

	sdklog "go.opentelemetry.io/otel/sdk/log"
)

type sophieLogger struct {
	logger         log.Logger
	loggerProvider *sdklog.LoggerProvider
}

func newSophieLogger(exporter sdklog.Exporter) *sophieLogger {

	processor := sdklog.NewBatchProcessor(exporter)

	logProvider := sdklog.NewLoggerProvider(
		sdklog.WithResource(config.Resource),
		sdklog.WithProcessor(processor),
	)

	logger := logProvider.Logger(config.ServiceName)

	return &sophieLogger{
		logger:         logger,
		loggerProvider: logProvider,
	}
}

func (l *sophieLogger) Debug(ctx context.Context, msg string) {
	record := log.Record{}

	record.SetBody(log.StringValue(msg))
	record.SetSeverity(5)
	record.SetSeverityText("debug")

	l.logger.Emit(
		ctx,
		record,
	)
}

func (l *sophieLogger) Info(ctx context.Context, msg string) {
	record := log.Record{}

	record.SetBody(log.StringValue(msg))
	record.SetSeverity(9)
	record.SetSeverityText("info")

	l.logger.Emit(
		ctx,
		record,
	)
}

func (l *sophieLogger) Warning(ctx context.Context, msg string) {
	record := log.Record{}

	record.SetBody(log.StringValue(msg))
	record.SetSeverity(13)
	record.SetSeverityText("warning")

	l.logger.Emit(
		ctx,
		record,
	)
}

func (l *sophieLogger) Error(ctx context.Context, msg string) {
	record := log.Record{}

	record.SetBody(log.StringValue(msg))
	record.SetSeverity(17)
	record.SetSeverityText("error")

	l.logger.Emit(
		ctx,
		record,
	)
}

func (l *sophieLogger) Fatal(ctx context.Context, msg string) {
	record := log.Record{}

	record.SetBody(log.StringValue(msg))
	record.SetSeverity(21)
	record.SetSeverityText("fatal")

	l.logger.Emit(
		ctx,
		record,
	)
}

func (l *sophieLogger) Close(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := l.loggerProvider.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
