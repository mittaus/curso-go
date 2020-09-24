package models

//Tabler permite cambiar el nombre de la tabla por defecto
type Tabler interface {
	TableName() string
}

// TableName sobreescribe la convención. Debió ser `alumns`, pero será `Alumno`
func (Alumno) TableName() string {
	return "Alumno"
}

// TableName de `modulos` a `Modulo`
func (Modulo) TableName() string {
	return "Modulo"
}

// TableName de `profesors` a `Profesor`
func (Profesor) TableName() string {
	return "Profesor"
}
