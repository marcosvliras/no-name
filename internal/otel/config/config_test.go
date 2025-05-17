package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	os.Setenv("SERVICE_NAME", "test-service")
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4317")

	fmt.Println(os.Getenv("SERVICE_NAME"))

	InitConfig()

	if Resource == nil {
		t.Errorf("Expected Resource to be set, but it is nil")
	}

	if Conn == nil {
		t.Errorf("Expected Conn to be set, but it is nil")
	}

	assert.Equal(t, "test-service", GetServiceName())
	assert.Equal(t, "localhost:4317", GetCollectorURL())

}
