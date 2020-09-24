package models

import "time"

// Alumno definición
type Alumno struct {
	ID              uint      `gorm:"column:Id" gorm:"primaryKey"`            //id
	DelegadoID      uint16    `gorm:"column:Delegado_Id;type:int" sql:"null"` //delegado_id
	Nombre          string    `gorm:"column:Nombre;type:varchar(50)"`         //nombre
	Apellido        string    `gorm:"column:Apellido;type:varchar(80)"`       //apellido
	FechaNacimiento time.Time `gorm:"column:FechaNacimiento;type:datetime"`   //fecha_nacimiento
	// Modulos         []Modulo `gorm:"many2many:AlumnoModulo;"`
	Modulos []Modulo `gorm:"many2many:AlumnoModulo;foreignKey:Id;joinForeignKey:AlumnoID;References:Codigo;JoinReferences:ModuloCodigo"`
}
