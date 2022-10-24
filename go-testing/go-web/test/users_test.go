package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/cmd/server/handler"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/internal/domain"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/internal/users"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// var response *httptest.ResponseRecorder = httptest.NewRecorder()

func createServer(mockStorage *mock.MockStorage) *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	repository := users.NewRepository(mockStorage)
	service := users.NewService(repository)
	controller := handler.NewUserController(service)

	server := gin.Default()

	usersEndpoint := server.Group("/users")
	// usersEndpoint.Use(handler.MiddlewareList()...)
	usersEndpoint.GET("", controller.GetAll())
	usersEndpoint.POST("", controller.CreateUser())
	usersEndpoint.PUT("/:id", handler.MiddlewareList(controller.Update())...)

	return server
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("TOKEN", "123456")

	return request, httptest.NewRecorder()
}

func TestUpdate(t *testing.T) {
	initialDatabase := []domain.User{{
		ID:       1,
		Name:     "Dionys",
		IsActive: true,
		Age:      25,
	}, {
		ID:       2,
		Name:     "Brian",
		IsActive: true,
		Age:      28,
	}, {
		ID:       3,
		Name:     "Matías",
		IsActive: true,
		Age:      25,
	}}
	mockStorage := &mock.MockStorage{
		DataMock: initialDatabase,
	}
	server := createServer(mockStorage)

	url := fmt.Sprintf("/users/%d", 3)
	request, responseRecorder := createRequestTest(http.MethodPut, url, `{"nombre": "Hernán", "estaActive": false, "edad": 25}`)

	var responseObject domain.User

	server.ServeHTTP(responseRecorder, request)
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &responseObject)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}
