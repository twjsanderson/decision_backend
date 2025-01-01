package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"

    "decision_backend/api/models"
)


func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
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
