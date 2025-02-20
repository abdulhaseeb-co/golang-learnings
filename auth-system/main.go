package main

import (
	"auth-system/config"
	"auth-system/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	config.InitDB()

	// Create new Gin router
	r := gin.Default()

	// Register routes
	routes.SetupRoutes(r)

	// Start server
	r.Run(":8080")
}
