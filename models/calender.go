package models

import "time"

type Calender struct {
	Date    time.Time `gorm:"primaryKey"`
	Weekday int
	Holiday bool
	Events  []Event // Calender has many Event
}
