package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/twjsanderson/decision_backend/internal/models"
)

func CreateUser(c *gin.Context) {
	var user models.User

	userClaims, err := c.Get("user")
	if err {
		fmt.Printf("An error occurred with Clerk Authentication: %v\n", err)
		c.Abort()
		return
	}
	if userClaims == nil {
		fmt.Printf("User not found: %v\n", userClaims)
		c.Abort()
		return
	}
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
	return
}

func GetUser(c *gin.Context) {
	// Your logic for fetching a user
	c.JSON(http.StatusOK, gin.H{"message": "User fetched"})
	return
}

func UpdateUser(c *gin.Context) {
	// Your logic for updating a user
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
	return
}

func DeleteUser(c *gin.Context) {
	// Your logic for deleting a user
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	return
}
