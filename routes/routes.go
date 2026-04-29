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
	r.HandleFunc("/api/personalidades", controllers.ExibePersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.ExibePersonalidadePorId).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
