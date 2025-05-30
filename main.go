package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GomdimApps/gmail-2fa/controllers"
	"github.com/GomdimApps/gmail-2fa/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, relying on environment variables")
	}

	database.ConnectDatabase()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	migrationsPath := os.Getenv("MIGRATIONS_PATH")

	migrateDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	if migrationsPath == "" {
		log.Fatal("MIGRATIONS_PATH environment variable not set.")
	}

	database.RunMigrations(migrateDSN, migrationsPath)

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	controllers.RegisterRoutes(router)

	// Start server
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080" // Default port
	}
	log.Printf("Server starting on port %s", serverPort)
	if err := router.Run(":" + serverPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
