package main

import (
	"github.com/BrunoPolaski/go-gin-api/models"
	"github.com/BrunoPolaski/go-gin-api/routes"
)

func main() {
	models.Students = []models.Student{
		{Name: "Bruno", CPF: "00000000000", RG: "00000000000"},
		{Name: "Pam", CPF: "22222222222", RG: "22222222222"},
	}
	routes.HandleRequests()
}
