package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/makhmudovazeez/effective_mobile/internal/router"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("cannot load .env file:", err)
		os.Exit(1)
	}
}

func main() {
	router.InitRoutes()
}
