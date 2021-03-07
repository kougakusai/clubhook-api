package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Number string
	Name   string
	email  string
	grade  uint
}
