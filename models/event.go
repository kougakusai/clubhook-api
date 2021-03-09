package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string
	Place        string
	OwnerID      uint      // foreign key
	CalenderDate time.Time // foreign key
	Owner        User      // Event belongs to User
	Calender     Calender  // Event belongs to Calender
	Vote         Vote      // Event has one Vote
	Users        []*User   `gorm:"many2many:event_users;"`
}
