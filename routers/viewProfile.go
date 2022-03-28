package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielnocce/projectx/db"
)

// Funcion visualizar el perfil de los usuarios
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {

		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
	}
	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro"+err.Error(), 400)

	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
