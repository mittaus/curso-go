package routes

import (
	"example.com/username/orm/controllers"
	"github.com/gorilla/mux"
)

// SetModuloRoutes agrega las rutas
func SetModuloRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/modulo/buscar", controllers.SearchModulos).Queries("q", "{q}").Methods("GET")
	subRouter.HandleFunc("/modulo", controllers.GetModulos).Methods("GET")
}
