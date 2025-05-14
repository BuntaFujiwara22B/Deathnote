package routes

import (
	"database/sql"
	"net/http"
	"encoding/json"
	"github.com/BuntaFujiwara22B/deathback/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Ruta principal con documentación básica de la API
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "Bienvenido a la API de Death Note",
			"endpoints": map[string]string{
				"create_victim":    "POST /victimas",
				"list_victims":     "GET /victimas",
				"get_victim":       "GET /victimas/{id}",
				"update_cause":     "PUT /victimas/{id}/cause",
				"update_details":   "PUT /victimas/{id}/details",
			},
			"rules": "Recuerda: Tienes 40 segundos para especificar la causa y 6m40s para detalles",
		}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	// Grupo de rutas para /victimas
	victimRouter := r.PathPrefix("/victimas").Subrouter()
	
	// Crear nueva víctima
	victimRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateVictima(db, w, r)
	}).Methods("POST")

	// Obtener todas las víctimas
	victimRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		controllers.ListVictimas(db, w, r)
	}).Methods("GET")

	// Obtener víctima específica
	victimRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetVictima(db, w, r)
	}).Methods("GET")

	// Actualizar causa de muerte
	victimRouter.HandleFunc("/{id}/cause", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateCause(db, w, r)
	}).Methods("PUT")

	// Actualizar detalles de muerte
	victimRouter.HandleFunc("/{id}/details", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateDetails(db, w, r)
	}).Methods("PUT")

	// Manejo de errores 404
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Endpoint no encontrado",
			"message": "Consulta / para ver los endpoints disponibles",
		})
	})

	return r
}