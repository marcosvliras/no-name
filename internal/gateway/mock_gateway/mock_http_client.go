package mock_gateway

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func GetHttpClientMock(t int8) *http.Client {
	bytes, err := os.ReadFile("mock_gateway/mock_stock.json")
	if err != nil {
		panic(err)
	}

	jsonContent := string(bytes)

	return &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) (*http.Response, error) {
			switch t {
			case 0:
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(jsonContent)),
					Header:     make(http.Header),
				}, nil
			case 1:
				return nil, fmt.Errorf("fake error")
			case 2:
				return &http.Response{
					StatusCode: 500,
					Body:       io.NopCloser(strings.NewReader("")),
					Header:     make(http.Header),
					Status:     "500 Internal Server Error",
				}, nil
			default:
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(jsonContent)),
					Header:     make(http.Header),
				}, nil
			}
		}),
	}
}
