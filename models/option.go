package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Name   string
	VoteID uint   // foreign key
	Vote   Vote   // Option belongs to Vote
	Casts  []Cast // Option has many Cast
}
