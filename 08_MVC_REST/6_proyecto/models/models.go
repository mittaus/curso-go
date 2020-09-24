package models

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

var db *gorm.DB

//Tabler permite cambiar el nombre de la tabla por defecto
type Tabler interface {
	TableName() string
}

// // TableName sobreescribe la convención. Debió ser `alumns`, pero será `Alumno`
// func (Alumno) TableName() string {
// 	return "Alumno"
// }

//Model hereda a todos
type Model struct {
	ID         int `gorm:"primaryKey" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// Setup initializes the database instance
func Setup() {
	// setting.DatabaseSetting.Type
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	var err error
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: setting.DatabaseSetting.TablePrefix, SingularTable: true},
	})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		defer sqlDB.Close()
	}
}

// // updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
// func updateTimeStampForCreateCallback(scope *gorm.DB) {
// 	if !scope.HasError() {
// 		nowTime := time.Now().Unix()
// 		if createTimeField, ok := scope. .FieldByName("CreatedOn"); ok {
// 			if createTimeField.IsBlank {
// 				createTimeField.Set(nowTime)
// 			}
// 		}

// 		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
// 			if modifyTimeField.IsBlank {
// 				modifyTimeField.Set(nowTime)
// 			}
// 		}
// 	}
// }

// // updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
// func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
// 	if _, ok := scope.Get("gorm:update_column"); !ok {
// 		scope.SetColumn("ModifiedOn", time.Now().Unix())
// 	}
// }

// // deleteCallback will set `DeletedOn` where deleting
// func deleteCallback(scope *gorm.Scope) {
// 	if !scope.HasError() {
// 		var extraOption string
// 		if str, ok := scope.Get("gorm:delete_option"); ok {
// 			extraOption = fmt.Sprint(str)
// 		}

// 		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

// 		if !scope.Search.Unscoped && hasDeletedOnField {
// 			scope.Raw(fmt.Sprintf(
// 				"UPDATE %v SET %v=%v%v%v",
// 				scope.QuotedTableName(),
// 				scope.Quote(deletedOnField.DBName),
// 				scope.AddToVars(time.Now().Unix()),
// 				addExtraSpaceIfExist(scope.CombinedConditionSql()),
// 				addExtraSpaceIfExist(extraOption),
// 			)).Exec()
// 		} else {
// 			scope.Raw(fmt.Sprintf(
// 				"DELETE FROM %v%v%v",
// 				scope.QuotedTableName(),
// 				addExtraSpaceIfExist(scope.CombinedConditionSql()),
// 				addExtraSpaceIfExist(extraOption),
// 			)).Exec()
// 		}
// 	}
// }

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
