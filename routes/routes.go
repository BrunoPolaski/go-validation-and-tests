package routes

import (
	"github.com/BrunoPolaski/go-gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./css")
	r.GET("/students", controllers.GetStudents)
	r.GET("/:name", controllers.Welcome)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.GetStudentById)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCpf)
	r.GET("/index", controllers.DisplayIndexPage)
	r.NoRoute(controllers.DisplayNotFound)
	r.Run()
}
