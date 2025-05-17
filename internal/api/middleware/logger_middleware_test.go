package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcosvliras/sophie/internal/api/middleware/mock_middleware"
	"github.com/marcosvliras/sophie/internal/otel/logging"
	"github.com/stretchr/testify/assert"
)

func TestLoggerMiddleware(t *testing.T) {

	err := logging.InitLogger("cli", nil)
	if err != nil {
		panic(err)
	}
	mockLogger := &mock_middleware.MockLogger{}
	logging.SLogger = mockLogger
	defer logging.SLogger.Close(context.Background())

	middleware := LoggerMiddleware()

	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.Use(middleware)
	server.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	// check request
	r := mockLogger.Infos[0]

	var loggedMsg map[string]interface{}
	_ = json.Unmarshal([]byte(r), &loggedMsg)

	assert.Equal(t, loggedMsg["HTTPMethod"], "GET")
	assert.Equal(t, loggedMsg["HTTPRoute"], "/test")
	assert.Equal(t, loggedMsg["HTTPStatus"], float64(200))
	assert.Equal(t, loggedMsg["QueryParams"], map[string]interface{}{})
	assert.Nil(t, loggedMsg["RequestBody"])

	// check response
	r2 := mockLogger.Infos[1]

	var loggedMsg2 map[string]interface{}
	_ = json.Unmarshal([]byte(r2), &loggedMsg2)

	assert.Equal(t, loggedMsg2["HTTPStatus"], float64(200))
	assert.Equal(t, loggedMsg2["ResponseBody"], "ok")

}
