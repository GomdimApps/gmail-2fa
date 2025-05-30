package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/GomdimApps/gmail-2fa/controllers"
	"github.com/GomdimApps/gmail-2fa/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	migrationsPath := strings.TrimSpace(os.Getenv("MIGRATIONS_PATH"))

	migrateDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))

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
		serverPort = "8080"
	}
	log.Printf("Server starting on port %s", serverPort)
	if err := router.Run(":" + serverPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
