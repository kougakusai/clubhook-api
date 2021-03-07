package models

import "time"

type Calender struct {
	Day     time.Time `gorm:"primaryKey"`
	Holiday bool
}
