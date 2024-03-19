package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoPolaski/go-gin-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func RouteTestingSetup() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestWelcomeStatusCode(t *testing.T) {
	r := RouteTestingSetup()
	r.GET("/:name", controllers.Welcome)
	req, _ := http.NewRequest("GET", "/Bruno", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "should be equal")
	responseMock := `{"API diz: ":"E aí Bruno"}`
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, responseMock, string(responseBody), "not equal")
}
