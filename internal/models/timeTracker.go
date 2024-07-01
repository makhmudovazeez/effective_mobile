package models

import "time"

type TimeTracker struct {
	ID     uint `gorm:"primaryKey"`
	Start  time.Time
	End    time.Time
	UserID uint
}

func (TimeTracker) TableName() string {
	return "time_trackers"
}
