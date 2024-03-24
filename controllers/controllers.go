package controllers

import (
	"net/http"

	"github.com/BrunoPolaski/go-gin-api/database"
	"github.com/BrunoPolaski/go-gin-api/models"
	"github.com/gin-gonic/gin"
)

// GetStudents godoc
// @Summary      Show all students in the database
// @Description  Get all students
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id path int true "Student ID"
// @Success      200  {object}  models.Student
// @Router       /students [get]
func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

// Welcome godoc
// @Summary      Show a nice welcome message with the name provided
// @Description  Get all students
// @Tags         name
// @Accept       json
// @Produce      json
// @Param        name path string true "Name"
// @Success      200  {object}  models.Student
// @Router       /{name} [get]
func Welcome(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API diz: ": "E aí " + name,
	})
}

// CreateStudent godoc
// @Summary      Create a new student
// @Description  Create a new student
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id path int true "Student ID"
// @Success      200  {object}  models.Student
// @Router       /students [post]
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: ": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

// GetStudentById	godoc
// @Summary      Show a student by ID
// @Description  Get a student by ID
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id path int true "Student ID"
// @Success      200  {object}  models.Student
// @Router       /students/{id} [get]
func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error: ": "Student Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: ": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func SearchStudentByCpf(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found: ": "student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DisplayIndexPage(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func DisplayNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{})
}
