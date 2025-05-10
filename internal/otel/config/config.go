package config

import "os"

var (
	ServiceName  = os.Getenv("SERVICE_NAME")
	CollectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
)

func InitConfig() {
	err := newResource()
	if err != nil {
		panic(err)
	}

	err = initConn()
	if err != nil {
		panic(err)
	}
}
