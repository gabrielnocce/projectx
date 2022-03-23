package middlew

import (
	"net/http"

	"github.com/gabrielnocce/projectx/db"
)

// Permite conocer el estado de la BD
func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if db.TestConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
