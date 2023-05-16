package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-auth/pkg/app"
	"go-auth/pkg/cache"
	"os"
)

func init() {
	// Load values in env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	server := app.New()

	// Init cache
	server.Cache = cache.New()

	server.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	server.SetupRoutes()

	server.Run(":3001")
}
