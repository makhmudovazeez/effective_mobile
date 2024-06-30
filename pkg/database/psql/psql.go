package psql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PSQLConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

func Connect(config PSQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tashkent",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
