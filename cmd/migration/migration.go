package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/makhmudovazeez/effective_mobile/internal/models"
	"github.com/makhmudovazeez/effective_mobile/pkg/database/psql"
	"os"
	"strconv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("cannot load .env file:", err)
		os.Exit(1)
	}
}

func main() {
	dbP, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("cannot parse db_port into int:", err)
		os.Exit(1)
	}
	config := psql.PSQLConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     dbP,
		DBName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	db, err := psql.Connect(config)
	if err != nil {
		fmt.Println("cannot connect to database:", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Println("error while migration users models:", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(&models.TimeTracker{}); err != nil {
		fmt.Println("error while migration time_trackers models:", err)
		os.Exit(1)
	}

	fmt.Println("successfully migrated")
}
