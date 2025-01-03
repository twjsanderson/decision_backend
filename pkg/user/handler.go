package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/twjsanderson/decision_backend/internal/auth"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func CreateUser(c *gin.Context) {
	var user models.User

	user, err := auth.ValidateClerkUser(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Request body check
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user fields"})
		return
	}

	// Success
	c.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
}

func GetUser(c *gin.Context) {
	// Your logic for fetching a user
	c.JSON(http.StatusOK, gin.H{"message": "User fetched"})
}

func UpdateUser(c *gin.Context) {
	// Your logic for updating a user
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func DeleteUser(c *gin.Context) {
	// Your logic for deleting a user
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
