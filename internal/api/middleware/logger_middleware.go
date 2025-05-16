package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/marcosvliras/sophie/internal/otel/logging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type requestInfo struct {
	HTTPMethod  string
	HTTPRoute   string
	HTTPStatus  int
	QueryParams url.Values
	RequestBody interface{}
}

type responseInfo struct {
	HTTPStatus   int
	ResponseBody interface{}
}

type customResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w customResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := otel.GetTextMapPropagator().Extract(
			c.Request.Context(),
			propagation.HeaderCarrier(c.Request.Header),
		)

		request := requestInfo{
			HTTPMethod:  c.Request.Method,
			HTTPRoute:   c.FullPath(),
			HTTPStatus:  c.Writer.Status(),
			QueryParams: c.Request.URL.Query(),
			RequestBody: c.Request.Body,
		}
		requestByte, err := json.MarshalIndent(request, "", "  ")
		if err != nil {
			logging.SLogger.Error(ctx, fmt.Sprintf("Error marshalling request: %v", err))
		}

		logging.SLogger.Info(ctx, string(requestByte))

		respBody := &bytes.Buffer{}
		cw := &customResponseWriter{body: respBody, ResponseWriter: c.Writer}
		c.Writer = cw

		c.Next()

		var body interface{}
		raw := respBody.String()

		if json.Valid([]byte(raw)) {
			_ = json.Unmarshal([]byte(raw), &body)
		} else {
			body = raw
		}

		response := responseInfo{
			HTTPStatus:   c.Writer.Status(),
			ResponseBody: body,
		}

		responseByte, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			logging.SLogger.Error(ctx, fmt.Sprintf("Error marshalling response: %v", err))
		}
		logging.SLogger.Info(ctx, string(responseByte))

	}
}
