package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/username/orm/models"
	"example.com/username/orm/utils"
)

// GetProfesores obtiene todos los alumnos
func GetProfesores(w http.ResponseWriter, r *http.Request) {
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
