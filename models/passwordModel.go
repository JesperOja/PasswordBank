package models

import "gorm.io/gorm"

type PasswordUnit struct {
	gorm.Model
	Site     string
	Password string
}
