package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name  string
	Users []User // Group has many User
}
