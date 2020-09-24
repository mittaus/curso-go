package utils

import (
	"fmt"

	"example.com/username/orm/models"
)

// MigrateDB migra la base de datos
func MigrateDB() {
	db := GetConnection()

	sqlDB, ok := db.DB()
	if ok != nil {
		defer sqlDB.Close()
	}

	fmt.Println("Migrating models....")
	// Automigrate se encarga de migrar la base de datos sí no se ha migrado, y lo hace a partir del modelo
	db.AutoMigrate(&models.Contact{}, &models.Alumno{}, &models.Modulo{}, &models.Profesor{})
}
