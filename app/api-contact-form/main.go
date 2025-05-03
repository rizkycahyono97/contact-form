package main

import (
	"api-contact-form/config"
	"api-contact-form/handlers"
	"api-contact-form/middleware"
	"api-contact-form/repositories"
	"api-contact-form/routes"
	"api-contact-form/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	//Load .env File
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	//Initialized Database
	config.InitDB()

	//Initialized Repository, Service, Handler
	contactRepository := repositories.NewContactRepository(config.DB)
	contactService := services.NewContactService(contactRepository)
	contactHandler := handlers.NewContactHandler(contactService)

	mainHandler := handlers.NewMainHandler()
	healthHandler := handlers.NewHealthHandler()

	// Create a new Gin router with default middleware (logger and recovery).
	router := gin.Default()

	// Apply the CORS middleware to the router.
	router.Use(cors.New(middleware.SetupCors()))

	// router
	routes.SetupRoutes(router, contactHandler, mainHandler, healthHandler)

	// Retrieve the application port from environment variables with a default value of "8080".
	appPort := config.GetEnv("APP_PORT", "8080")

	// run application
	if err := router.Run(":" + appPort); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
