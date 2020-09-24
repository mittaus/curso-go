package routes

import (
	"example.com/username/orm/controllers"
	"github.com/gorilla/mux"
)

// SetAlumnosRoutes agrega las rutas de contactos
func SetAlumnosRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/alumno/buscar", controllers.SearchAlumnos).Queries("q", "{q}").Methods("GET")
	subRouter.HandleFunc("/alumno/{id}", controllers.GetAlumnoByID).Methods("GET")
	subRouter.HandleFunc("/alumno", controllers.GetAlumnos).Methods("GET")
	subRouter.HandleFunc("/alumno", controllers.StoreAlumno).Methods("POST")
	subRouter.HandleFunc("/alumno/{id}", controllers.UpdateAlumno).Methods("PUT")
	subRouter.HandleFunc("/alumno/{id}", controllers.DeleteAlumno).Methods("DELETE")
}
