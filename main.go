package main

import (
	"fmt"

	"github.com/MatheusVictorMartins/goRestAPI/db"
	"github.com/MatheusVictorMartins/goRestAPI/routes"
)

func main() {

	fmt.Println("Starting Go server...")
	fmt.Println("Connecting to database...")
	db.ConectBD()
	routes.HandleRequest()
}
