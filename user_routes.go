package routes

import (
	"github.com/gin-gonic/gin"

	"9projectstructure/controllers"
)

func UserRoutes(router *gin.Engine, userController *controllers.UserController) {

	router.POST("/users", userController.CreateUser)

	router.GET("/users", userController.GetUsers)

	router.GET("/users/:id", userController.GetUser)

	router.PUT("/users/:id", userController.UpdateUser)

	router.DELETE("/users/:id", userController.DeleteUser)
}
