package routes

import (
	"log"
	"net/http"

	"github.com/MatheusVictorMartins/goRestAPI/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.ExibePersonalidades)

	log.Fatal(http.ListenAndServe(":8000", r))
}
