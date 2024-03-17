package routes

import (
	"github.com/BrunoPolaski/go-gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetStudents)
	r.GET("/:name", controllers.Welcome)
	r.Run()
}
