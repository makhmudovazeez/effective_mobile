package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/makhmudovazeez/effective_mobile/internal/handlers"
	"github.com/makhmudovazeez/effective_mobile/pkg/database/psql"
	"os"
	"strconv"
)

func InitRoutes() {
	appPort := os.Getenv("APP_PORT")
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

	userHandler := handlers.NewUserHandler(db)
	timeTrackerHandler := handlers.NewTimeTrackerHandler(db)

	_ = userHandler
	_ = timeTrackerHandler

	r := gin.Default()

	r.Run(fmt.Sprintf(":%s", appPort))
}
