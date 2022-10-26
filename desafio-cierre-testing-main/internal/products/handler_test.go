package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type serviceMock struct {
	getAllBySellerWasCalled bool
	repositoryError         error
}

func (ms *serviceMock) GetAllBySeller(sellerID string) ([]Product, error) {
	if ms.repositoryError != nil {
		return nil, ms.repositoryError
	}

	ms.getAllBySellerWasCalled = true
	return []Product{{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	}}, nil
}

func createServer(mockService *serviceMock) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	controller := NewHandler(mockService)

	endpoint := server.Group("/products")
	endpoint.GET("", controller.GetProducts)

	return server
}

func createTestRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	request.Header.Add("Content-Type", "application/json")

	return request, httptest.NewRecorder()
}

func TestGetProductsStautsOK(t *testing.T) {
	productsExpected := []Product{{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	}}

	serviceMock := &serviceMock{
		getAllBySellerWasCalled: false,
	}

	server := createServer(serviceMock)

	url := fmt.Sprintf("/products?%v=%v", "seller_id", "mock")
	request, response := createTestRequest(http.MethodGet, url, "")

	server.ServeHTTP(response, request)

	var responseObject []Product
	err := json.Unmarshal(response.Body.Bytes(), &responseObject)

	assert.True(t, serviceMock.getAllBySellerWasCalled)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	assert.True(t, len(responseObject) > 0)
	assert.Equal(t, productsExpected, responseObject)
}

func TestGetProductsStautsInternalServerError(t *testing.T) {
	errorExpected := map[string]string{
		"error": "something went wrong",
	}

	serviceMock := &serviceMock{
		getAllBySellerWasCalled: false,
		repositoryError:         errors.New("something went wrong"),
	}

	server := createServer(serviceMock)

	url := fmt.Sprintf("/products?%v=%v", "seller_id", "mock")
	request, response := createTestRequest(http.MethodGet, url, "")

	server.ServeHTTP(response, request)

	var responseMap map[string]string
	err := json.Unmarshal(response.Body.Bytes(), &responseMap)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, errorExpected, responseMap)
}
