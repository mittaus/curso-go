package models

// Profesor definition
type Profesor struct {
	RFC          uint   `gorm:"primaryKey;autoIncrement:false"`
	Nombres      string `gorm:"type:varchar(100)"`
	Direccion    string `gorm:"type:varchar(200)"`
	Telefono     string `gorm:"type:char(9)"`
	ModuloCodigo uint   `gorm:"column:Modulo_Codigo;type:int"`
}
