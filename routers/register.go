package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielnocce/projectx/db"
	"github.com/gabrielnocce/projectx/models"
)

// Funcion para crear el registro en mongo db
func Register(w http.ResponseWriter, r *http.Request) {

	var usr models.User
	err := json.NewDecoder(r.Body).Decode(&usr) // body es un objeto Stream osea se recibe y se elimina de memoria automaticamente

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(usr.Email) == 0 {

		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(usr.Password) < 6 {
		http.Error(w, "La contraseña del usuario debe ser de al menos 6 caracteres", 400)
		return

	}

	_, encontrado, _ := db.DuplicateControl(usr.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario creado con ese email", 400)
		return
	}

	_, status, err := db.InsertUsr(usr)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro del usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro registrar el usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
