package models

// Modulo definition
type Modulo struct {
	Codigo   uint       `gorm:"primaryKey"`
	Nombre   string     `gorm:"type:varchar(20)"`
	Alumnos  []Alumno   `gorm:"many2many:AlumnoModulo;foreignKey:Codigo;joinForeignKey:ModuloCodigo;References:Id;JoinReferences:Alumno_Id"`
	Profesor []Profesor `gorm:"foreignKey:RFC;references:Codigo"`
}

// Resultado de join Modulo y Alumno
type ResultadoModuloAlumno struct {
	Modulo Modulo `gorm:"embedded"`
	Alumno Alumno `gorm:"embedded"`
}
