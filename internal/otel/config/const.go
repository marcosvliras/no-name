package config

import "os"

func GetServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func GetCollectorURL() string {
	return os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
}
