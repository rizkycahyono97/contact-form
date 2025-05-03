package routes

import (
	"api-contact-form/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, contactHandler *handlers.ContactHandler, mainHandler *handlers.MainHandler, healthHandler *handlers.HealthHandler) {
	// Define application routes.
	router.GET("/", mainHandler.MainHandler)
	router.GET("/health", healthHandler.HealthCheck)
	router.GET("/contacts", contactHandler.GetContacts)
	router.GET("/contacts/:id", contactHandler.GetContact)
	router.POST("/contacts", contactHandler.CreateContact)
	router.PUT("/contacts/:id", contactHandler.UpdateContact)
	router.DELETE("/contacts/:id", contactHandler.DeleteContact)
}
