package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoPolaski/go-gin-api/controllers"
	"github.com/gin-gonic/gin"
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
	if response.Code != http.StatusOK {
		t.Fatalf("Status error: received value was %d, must be %d", response.Code, http.StatusOK)
	}
}
