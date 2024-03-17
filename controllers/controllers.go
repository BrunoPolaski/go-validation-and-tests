package controllers

import (
	"github.com/BrunoPolaski/go-gin-api/models"
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Welcome(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API diz: ": "E aí " + name,
	})
}
