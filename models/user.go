package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string
	Email   string
	Grade   uint
	GroupID uint     // foreign key
	Group   Group    `gorm:"save_association:false"` // User belongs to Group
	Casts   []Cast   // User has many Cast
	Events  []*Event `gorm:"many2many:event_users;"`
}
