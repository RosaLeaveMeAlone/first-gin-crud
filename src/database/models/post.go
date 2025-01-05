package models

import "gorm.io/gorm"

// Post representa el modelo en la base de datos.
type Post struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);not null"`
	Body  string `gorm:"type:text;not null"`
}
