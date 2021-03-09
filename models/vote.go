package models

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	Name    string
	EventID uint     // foreign key
	Options []Option // Vote has many Option
}
