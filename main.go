package main

import (
	"fmt"

	"github.com/MatheusVictorMartins/goRestAPI/models"
	"github.com/MatheusVictorMartins/goRestAPI/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Descricao: "Descricao 1"},
		{Nome: "Nome 2", Descricao: "Descricao 2"},
	}
	fmt.Println("Starting Go server...")
	routes.HandleRequest()
}
