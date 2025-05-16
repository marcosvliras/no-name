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
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"

	"go.opentelemetry.io/otel/log"

	sdklog "go.opentelemetry.io/otel/sdk/log"
)

var SLogger *SophieLogger

func InitLogger() error {
	var err error

	ctx := context.Background()

	SLogger, err = NewSophieLogger(ctx)
	if err != nil {
		return err
	}
	return nil

}

func logExporter(ctx context.Context) (*otlploggrpc.Exporter, error) {

	exporter, err := otlploggrpc.New(
		ctx,
		otlploggrpc.WithGRPCConn(config.Conn),
		otlploggrpc.WithInsecure(),
	)

	if err != nil {
		return nil, err
	}
	return exporter, nil
}

//func stdoutExporter(ctx context.Context) (*stdoutlog.Exporter, error) {
//	exporter, err := stdoutlog.New(stdoutlog.WithPrettyPrint())
//	if err != nil {
//		return nil, err
//	}
//	return exporter, nil
//}

type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warning(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

type SophieLogger struct {
	logger         log.Logger
	loggerProvider *sdklog.LoggerProvider
}

func NewSophieLogger(ctx context.Context) (*SophieLogger, error) {
	otelExporter, err := logExporter(ctx)
	if err != nil {
		return nil, err
	}

	processor := sdklog.NewBatchProcessor(otelExporter)

	logProvider := sdklog.NewLoggerProvider(
		sdklog.WithResource(config.Resource),
		sdklog.WithProcessor(processor),
	)

	logger := logProvider.Logger(config.ServiceName)

	return &SophieLogger{
		logger:         logger,
		loggerProvider: logProvider,
	}, err
}

func (l *SophieLogger) Debug(ctx context.Context, msg string) {
	record := log.Record{}

	record.SetBody(log.StringValue(msg))
	record.SetSeverity(5)
	record.SetSeverityText("debug")

	l.logger.Emit(
		ctx,
		record,
	)
}

func (l *SophieLogger) Info(ctx context.Context, msg string) {
	record := log.Record{}

	record.SetBody(log.StringValue(msg))
	record.SetSeverity(9)
	record.SetSeverityText("info")

	l.logger.Emit(
		ctx,
		record,
	)
}

func (l *SophieLogger) Warning(ctx context.Context, msg string) {
	record := log.Record{}

	record.SetBody(log.StringValue(msg))
	record.SetSeverity(13)
	record.SetSeverityText("warning")

	l.logger.Emit(
		ctx,
		record,
	)
}

func (l *SophieLogger) Error(ctx context.Context, msg string) {
	record := log.Record{}

	record.SetBody(log.StringValue(msg))
	record.SetSeverity(17)
	record.SetSeverityText("error")

	l.logger.Emit(
		ctx,
		record,
	)
}

func (l *SophieLogger) Fatal(ctx context.Context, msg string) {
	record := log.Record{}

	record.SetBody(log.StringValue(msg))
	record.SetSeverity(21)
	record.SetSeverityText("fatal")

	l.logger.Emit(
		ctx,
		record,
	)
}

func (l *SophieLogger) Close(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := l.loggerProvider.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
