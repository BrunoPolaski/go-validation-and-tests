package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/BrunoPolaski/go-gin-api/controllers"
	"github.com/BrunoPolaski/go-gin-api/database"
	"github.com/BrunoPolaski/go-gin-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func RouteTestingSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateMockStudent() {
	aluno := models.Student{Name: "Mock", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeleteMockStudent() {
	database.DB.Delete(&models.Student{}, ID)
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

func TestGetStudents(t *testing.T) {
	database.ConnectToDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()
	r := RouteTestingSetup()
	r.GET("/students", controllers.GetStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "should be equal")
}

func TestSearchStudentByCpf(t *testing.T) {
	database.ConnectToDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()
	r := RouteTestingSetup()
	r.GET("/students/:cpf", controllers.SearchStudentByCpf)
	req, _ := http.NewRequest("GET", "/students/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "should be equal")
}

func TestGetStudentById(t *testing.T) {
	database.ConnectToDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()
	r := RouteTestingSetup()
	r.GET("/students/:id", controllers.GetStudentById)
	req, _ := http.NewRequest("GET", "/students/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var mockStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &mockStudent)
	fmt.Println(mockStudent.Name)
	assert.Equal(t, "Mock", mockStudent.Name, "should be equal")
	assert.Equal(t, "12345678901", mockStudent.CPF, "should be equal")
	assert.Equal(t, "123456789", mockStudent.RG, "should be equal")
	assert.Equal(t, http.StatusOK, response.Code, "should be equal")
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectToDatabase()
	CreateMockStudent()
	r := RouteTestingSetup()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	req, _ := http.NewRequest("DELETE", "/students/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "should be equal")
}

func TestEditStudent(t *testing.T) {
	database.ConnectToDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()
	r := RouteTestingSetup()
	r.PATCH("/students/:id", controllers.UpdateStudent)
	student := models.Student{Name: "Mock Updated", CPF: "98340923483", RG: "987654321"}
	jsonStudent, _ := json.Marshal(student)
	req, _ := http.NewRequest("PATCH", "/students/"+strconv.Itoa(ID), bytes.NewBuffer(jsonStudent))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var updatedMockStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &updatedMockStudent)
	assert.Equal(t, "Mock Updated", updatedMockStudent.Name, "should be equal")
	assert.Equal(t, "98340923483", updatedMockStudent.CPF, "should be equal")
	assert.Equal(t, "987654321", updatedMockStudent.RG, "should be equal")
}
