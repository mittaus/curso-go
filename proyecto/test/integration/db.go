package integration

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
		DbUser,
		DbPassword,
		DbHost,
		DbName)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	// &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{TablePrefix: setting.DatabaseSetting.TablePrefix, SingularTable: true},
	// })

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	// sqlDB, err := db.DB()
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	return db, err
}

//closes the  database connection
// func CloseDB() {
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		defer sqlDB.Close()
// 	}
// }
