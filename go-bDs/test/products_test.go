package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bootcamp-go/go-bDs/cmd/server/routes"
	"github.com/bootcamp-go/go-bDs/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	os.Setenv("DBNAME", "storage")
	os.Setenv("DBPASS", "Meli_Sprint#123")
	os.Setenv("DBUSER", "meli_sprint_user")

	webEngine, db := db.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)

	router := routes.NewRouter(webEngine, db)

	router.MapRoutes()

	return webEngine
}

func createRquest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	request.Header.Set("Content-Type", "application/json")

	return request, httptest.NewRecorder()
}

func TestGetByName_StatusOK(t *testing.T) {
	request, response := createRquest(http.MethodGet, "/api/v1/products/lampara", "")

	webEngine := createServer()
	webEngine.ServeHTTP(response, request)

	t.Log(response.Body.String())
	assert.Equal(t, http.StatusOK, response.Code)
}
