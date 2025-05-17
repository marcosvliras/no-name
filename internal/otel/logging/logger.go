package logging

import (
	"context"
	"fmt"

	sdklog "go.opentelemetry.io/otel/sdk/log"
)

type Logger interface {
	Debug(ctx context.Context, msg string)
	Info(ctx context.Context, msg string)
	Warning(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
	Fatal(ctx context.Context, msg string)
	Close(ctx context.Context) error
}

var SLogger Logger

// logType should be "cli" or "api"
// If logType is "api", an exporter must be provided
func InitLogger(logType string, exporter sdklog.Exporter) error {

	if logType == "cli" {
		SLogger = NewCliLogger()
		return nil
	} else if logType == "api" && exporter != nil {
		SLogger = newSophieLogger(exporter)
		return nil
	}

	return fmt.Errorf("invalid log type: %s or exporter is nil", logType)

}
