package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/username/orm/models"
	"example.com/username/orm/utils"
	"github.com/gorilla/mux"
)

// SearchAlumnos obtiene todos los alumnos
func SearchAlumnos(w http.ResponseWriter, r *http.Request) {
	texto := mux.Vars(r)["q"]

	alumnos := []models.Alumno{}
	db := utils.GetConnection()
	sqlDB, ok := db.DB()
	if ok != nil {
		defer sqlDB.Close()
	}

	db.Where("Nombre LIKE ?", "%"+texto+"%").Find(&alumnos)
	j, _ := json.Marshal(alumnos)
	utils.SendResponse(w, http.StatusOK, j)
}

// GetAlumnoByID obtiene un alumno por su ID
func GetAlumnoByID(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	db := utils.GetConnection()
	sqlDB, ok := db.DB()
	if ok != nil {
		defer sqlDB.Close()
	}

	alumno := models.Alumno{}
	resultado := db.Find(&alumno, id)

	if resultado.RowsAffected > int64(0) {
		j, _ := json.Marshal(alumno)
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		utils.SendErr(w, http.StatusNotFound)
	}
}

// GetAlumnos obtiene todos los alumnos
func GetAlumnos(w http.ResponseWriter, r *http.Request) {
	alumnos := []models.Alumno{}
	db := utils.GetConnection()
	sqlDB, ok := db.DB()
	if ok != nil {
		defer sqlDB.Close()
	}

	db.Find(&alumnos) //SELECT * FROM Alumno
	j, _ := json.Marshal(alumnos)
	utils.SendResponse(w, http.StatusOK, j)
}

// StoreAlumno guarda un nuevo alumno
func StoreAlumno(w http.ResponseWriter, r *http.Request) {

	alumno := models.Alumno{}

	db := utils.GetConnection()
	sqlDB, ok := db.DB()
	if ok != nil {
		defer sqlDB.Close()
	}

	err := json.NewDecoder(r.Body).Decode(&alumno)
	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}

	err = db.Create(&alumno).Error
	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusInternalServerError)
		return
	}

	j, _ := json.Marshal(alumno)
	utils.SendResponse(w, http.StatusCreated, j)
}

// UpdateAlumno modifica los datos de un alumno por su ID
func UpdateAlumno(w http.ResponseWriter, r *http.Request) {

	alumnoActual := models.Alumno{}
	alumnoCambios := models.Alumno{}

	id := mux.Vars(r)["id"]

	db := utils.GetConnection()
	sqlDB, ok := db.DB()
	if ok != nil {
		defer sqlDB.Close()
	}

	db.Find(&alumnoActual, id)
	if alumnoActual.ID > 0 {

		err := json.NewDecoder(r.Body).Decode(&alumnoCambios)
		if err != nil {

			utils.SendErr(w, http.StatusBadRequest)
			return
		}

		db.Model(&alumnoActual).Updates(alumnoCambios)

		j, _ := json.Marshal(alumnoActual)
		utils.SendResponse(w, http.StatusOK, j)
	} else {

		utils.SendErr(w, http.StatusNotFound)
	}
}

// DeleteAlumno elimina un alumno por ID
func DeleteAlumno(w http.ResponseWriter, r *http.Request) {

	alumno := models.Alumno{}

	id := mux.Vars(r)["id"]

	db := utils.GetConnection()
	sqlDB, ok := db.DB()
	if ok != nil {
		defer sqlDB.Close()
	}

	db.Find(&alumno, id)
	if alumno.ID > 0 {

		db.Delete(alumno)
		utils.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {

		utils.SendErr(w, http.StatusNotFound)
	}
}
