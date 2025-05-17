package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/marcosvliras/sophie/internal/api/middleware/mock_middleware"
)

func TestRequestCounterMiddleware(t *testing.T) {

	ctrl := gomock.NewController(t)
	counter := mock_middleware.NewMockCounter(ctrl)

	// Expect the Add method to be called 5 times with any context, an increment of 1,
	counter.
		EXPECT().
		Add(gomock.Any(), int64(1), gomock.Any()).
		Times(5)

	middleware := RequestCounterMiddleware(counter)

	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.Use(middleware)
	server.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	for i := 0; i < 5; i++ {
		server.ServeHTTP(w, req)
	}

}
