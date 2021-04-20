package models

import "gorm.io/gorm"

type Cast struct {
	gorm.Model
	OptionID uint   // foreign key
	UserID   uint   // foreign key
	Option   Option // Cast belongs to Option
	User     User   // Cast belongs to User
}
