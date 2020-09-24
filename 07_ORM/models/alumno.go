package models

import "time"

// Alumno definición
type Alumno struct {
	ID              uint      `gorm:"column:Id" gorm:"primaryKey"`
	DelegadoID      uint16    `gorm:"column:Delegado_Id;type:int" sql:"null"`
	Nombre          string    `gorm:"column:Nombre;type:varchar(50)"`
	Apellido        string    `gorm:"column:Apellido;type:varchar(80)"`
	FechaNacimiento time.Time `gorm:"column:FechaNacimiento;type:datetime"`
	// Modulos         []Modulo `gorm:"many2many:AlumnoModulo;"`
	Modulos []Modulo `gorm:"many2many:AlumnoModulo;foreignKey:Id;joinForeignKey:AlumnoID;References:Codigo;JoinReferences:ModuloCodigo"`
}
