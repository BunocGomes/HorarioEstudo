package models

import "gorm.io/gorm"

// Materia representa a matéria de estudo no banco de dados
type Materia struct {
	gorm.Model
	Nome      string `gorm:"unique;not null"`
	Professor string
	Descricao string
}