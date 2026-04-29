package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MatheusVictorMartins/goRestAPI/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Um teste bom")
}

func ExibePersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
