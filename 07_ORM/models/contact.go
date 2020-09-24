package models

import "github.com/jinzhu/gorm"

// Contact modelo para contactos
type Contact struct {
	gorm.Model
	Nombre      string `json:"nombre" gorm:"type:varchar(35)"`
	Edad        uint   `json:"edad" gorm:"type:int"`
	Telefono    string `json:"telefono" gorm:"type:varchar(20)"`
	Direccion   string `json:"direccion" gorm:"type:varchar(50)"`
	Email       string `json:"email" gorm:"type:varchar(30)"`
	Descripcion string `json:"descripcion" gorm:"type:TEXT"`
}
