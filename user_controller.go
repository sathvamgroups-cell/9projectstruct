package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"9projectstructure/models"
)

type UserController struct {
	DB *gorm.DB
}

// CREATE USER
func (uc *UserController) CreateUser(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})

		return
	}

	result := uc.DB.Create(&user)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusCreated, user)
}

// GET ALL USERS
func (uc *UserController) GetUsers(c *gin.Context) {

	var users []models.User

	result := uc.DB.Find(&users)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})

		return
	}

	c.JSON(http.StatusOK, users)
}

// GET SINGLE USER
func (uc *UserController) GetUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	result := uc.DB.First(&user, "id = ?", id)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

// UPDATE USER
func (uc *UserController) UpdateUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	// Find existing user
	result := uc.DB.First(&user, "id = ?", id)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})

		return
	}

	// Receive new data
	var input models.User

	err := c.ShouldBindJSON(&input)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})

		return
	}

	// Update user information
	user.Name = input.Name

	// Save changes
	result = uc.DB.Save(&user)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

// DELETE USER
func (uc *UserController) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	// Find existing user
	result := uc.DB.First(&user, "id = ?", id)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})

		return
	}

	// Delete user
	result = uc.DB.Delete(&user)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
