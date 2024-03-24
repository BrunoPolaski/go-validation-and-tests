package main

import (
	"github.com/BrunoPolaski/go-gin-api/database"
	"github.com/BrunoPolaski/go-gin-api/routes"
)

// @title           Student API
// @version         1.0
// @description     A simple API to manage students
// @termsOfService  http://swagger.io/terms/

// @contact.name   Bruno Polaski
// @contact.email  polaskibruno03@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
