package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/marcosvliras/no-name/internal/service/mock_service"
	"github.com/marcosvliras/no-name/stock"
	"github.com/stretchr/testify/assert"
)

func TestStocksCtrl(t *testing.T) {

	t.Run("Test when correct symbol is passed", func(t *testing.T) {
		gomockCtrl := gomock.NewController(t)

		svc := mock_service.NewMockISVC(gomockCtrl)

		actualPrice := 9.95
		maxPrice := 7.27
		svc.
			EXPECT().
			GetStockData(gomock.Any()).
			Return([]stock.AggStockData{
				{
					Stock:         "ITSA4.SAO",
					ActualPrice:   &actualPrice,
					MaxStockPrice: &maxPrice,
				},
			})

		ctrl := StocksCtrl{
			svc: svc,
		}
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/stocks", ctrl.Handle)

		req, _ := http.NewRequest(http.MethodGet, "/stocks?symbol=BBAS3", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		expectedJSON := `[
			{
				"Stock": "ITSA4.SAO",
				"MaxStockPrice": 7.27,
				"ActualPrice": 9.95
			}
		]`

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, expectedJSON, w.Body.String())

	})

	t.Run("Test when no symbols provided", func(t *testing.T) {
		gomockCtrl := gomock.NewController(t)

		svc := mock_service.NewMockISVC(gomockCtrl)

		ctrl := NewStocksCtrl(svc)
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/stocks", ctrl.Handle)

		req, _ := http.NewRequest(http.MethodGet, "/stocks", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		expectedJSON := `{
		  		"error": "No symbols provided"
		}`

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, expectedJSON, w.Body.String())
	})
}
