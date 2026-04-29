package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MatheusVictorMartins/goRestAPI/db"
	"github.com/MatheusVictorMartins/goRestAPI/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func ExibePersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	db.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ExibePersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	db.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CriaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	db.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}
