package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ioiiiomio/Golang/models"
	"github.com/ioiiiomio/Golang/routes"
	"github.com/ioiiiomio/Golang2024/config"
	//so I should specify current file
	"log"
)

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// DB connection
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Routes
	routes.SetupRoutes(r)

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
