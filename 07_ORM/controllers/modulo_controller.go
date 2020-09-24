package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/username/orm/models"
	"example.com/username/orm/utils"
	"github.com/gorilla/mux"
)

// SearchModulos obtiene todos los modulos
func SearchModulos(w http.ResponseWriter, r *http.Request) {
	texto := mux.Vars(r)["q"]

	// modulo := []models.Modulo{}
	resultado := []models.ResultadoModuloAlumno{}
	db := utils.GetConnection()
	sqlDB, ok := db.DB()
	if ok != nil {
		defer sqlDB.Close()
	}

	db.Raw("spu_modulosSearch ?,?", "%"+texto+"%", 1).Scan(&resultado)

	j, _ := json.Marshal(resultado)
	utils.SendResponse(w, http.StatusOK, j)
}

// GetModulos obtiene todos los alumnos
func GetModulos(w http.ResponseWriter, r *http.Request) {
	alumnos := []models.Alumno{}
	db := utils.GetConnection()
	sqlDB, ok := db.DB()
	if ok != nil {
		defer sqlDB.Close()
	}

	db.Find(&alumnos)
	j, _ := json.Marshal(alumnos)
	utils.SendResponse(w, http.StatusOK, j)
}
