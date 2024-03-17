package main

import (
	"github.com/BrunoPolaski/go-gin-api/database"
	"github.com/BrunoPolaski/go-gin-api/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
