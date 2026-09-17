package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"9projectstructure/config"
	"9projectstructure/controllers"
	"9projectstructure/models"
	"9projectstructure/routes"
)

func main() {

	// Connect to database
	db := config.ConnectDatabase()

	// Create/update database table
	db.AutoMigrate(&models.User{})

	// Create Gin router
	router := gin.Default()

	// Home route
	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to my Gin API",
		})

	})

	// Create User Controller
	userController := controllers.UserController{
		DB: db,
	}

	// Register User Routes
	routes.UserRoutes(router, &userController)

	// Start server
	router.Run(":4000")
}
