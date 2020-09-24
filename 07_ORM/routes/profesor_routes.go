package routes

import (
	"example.com/username/orm/controllers"
	"github.com/gorilla/mux"
)

// SetProfesorRoutes agrega las rutas
func SetProfesorRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/profesor", controllers.GetProfesores).Methods("GET")
}
