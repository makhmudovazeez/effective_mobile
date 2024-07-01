package models

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey"`
	PassportSeries string    `gorm:"size:10"` //To accept series whatever citizen length is set to 10
	PassportNumber string    `gorm:"size:10"` //The same situation with number
	Surname        string    `gorm:"size:256"`
	Name           string    `gorm:"size:256"`
	Patronymic     string    `gorm:"size:256"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	TimeTrackers []TimeTracker
}

func (User) TableName() string {
	return "users"
}
