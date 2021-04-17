package models

import (
	"time"

	"gorm.io/gorm"
)

type Calender struct {
	Date      time.Time `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt
	Weekday   int
	Holiday   bool
	Events    []Event // Calender has many Event
}
