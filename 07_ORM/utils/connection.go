package utils

import (
	"log"

	// "github.com/jinzhu/gorm"

	// sqlserver
	// _ "github.com/denisenkom/go-mssqldb"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// GetConnection obtiene una conexión a la base de datos
func GetConnection() *gorm.DB {
	dsn := "sqlserver://sa:Pa$$w0rd@localhost:1433?database=CONTACTOS_DEV&connection+timeout=30"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
