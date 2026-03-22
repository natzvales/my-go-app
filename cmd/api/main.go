package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/natz/go-lib-app/internal/config"
	"github.com/natz/go-lib-app/internal/container"
	"github.com/natz/go-lib-app/internal/database"
	"github.com/natz/go-lib-app/internal/server"

	_ "github.com/natz/go-lib-app/internal/modules/auth"
	_ "github.com/natz/go-lib-app/internal/modules/books"
)

func main() {

	//load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Connect to DB
	db := database.Connect()
	cfg := config.LoadConfig()

	container := container.NewContainer(db, cfg)

	// Collect modules

	modules := server.LoadModules(container)

	//Start the server

	router := server.NewServer(modules)
	// log.Println("Starting server on :8080")
	fmt.Printf("Starting server on :%s\n", port)
	router.Run(":" + port)
}
