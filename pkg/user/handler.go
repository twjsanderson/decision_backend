package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/twjsanderson/decision_backend/internal/auth"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func ValidateRequest(c *gin.Context) (*models.User, error) {
	var requestBody models.User

	// Get the route path from the context
	routePath := c.FullPath()

	// get user_id (as id) from params
	userId := c.Query("id")
	if userId == "" {
		return nil, fmt.Errorf("missing user id")
	}
	requestBody.Id = userId // assign userId from params to requestBody

	if routePath != "/user/get" && routePath != "/user/delete" {
		// Validate the JSON format for all req. except GET & DELETE
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			return nil, fmt.Errorf("invalid request body format")
		}
	}

	if routePath == "/user/create" {
		if requestBody.Id == "" {
			return nil, fmt.Errorf("missing or empty required field(s)")
		}
		if requestBody.Email == "" ||
			requestBody.FirstName == nil ||
			*requestBody.FirstName == "" ||
			requestBody.LastName == nil ||
			*requestBody.LastName == "" {
			return nil, fmt.Errorf("missing or empty required field(s)")
		}
		if requestBody.IsAdmin {
			return nil, fmt.Errorf("unauthorized field isAdmin found")
		}
	}

	// Return the valid request
	return &requestBody, nil
}

func CreateUser(c *gin.Context) {
	// Authenticate
	clerkUser, authenticatedErr := auth.AuthenticateClerkUser(c)
	if authenticatedErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": authenticatedErr.Error()})
		return
	}

	// Validate Request Body
	validatedRequestBody, validatedRequestErr := ValidateRequest(c)
	if validatedRequestErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validatedRequestErr.Error()})
		return
	}
	status, err := CreateUserService(&clerkUser, validatedRequestBody)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	// Success
	c.JSON(status, gin.H{"message": "user created"})
}

func GetUser(c *gin.Context) {
	// Authenticate
	clerkUser, authenticatedErr := auth.AuthenticateClerkUser(c)
	if authenticatedErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": authenticatedErr.Error()})
		return
	}
	// Validate Request Body
	validatedRequestBody, validatedRequestErr := ValidateRequest(c)
	if validatedRequestErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validatedRequestErr.Error()})
		return
	}

	dbUser, dbStatus, dbErr := GetUserService(&clerkUser, validatedRequestBody)
	if dbErr != nil {
		c.JSON(dbStatus, gin.H{"error": dbErr.Error()})
		return
	}

	// Success
	c.JSON(dbStatus, gin.H{"data": dbUser})
}

func GetAllUsers(c *gin.Context) {
	// get all users
}

func UpdateUser(c *gin.Context) {
	// Authenticate
	clerkUser, authenticatedErr := auth.AuthenticateClerkUser(c)
	if authenticatedErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": authenticatedErr.Error()})
		return
	}
	// Validate Request Body
	validatedRequestBody, validatedRequestErr := ValidateRequest(c)
	if validatedRequestErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validatedRequestErr.Error()})
		return
	}

	dbUser, dbStatus, dbErr := UpdateUserService(&clerkUser, validatedRequestBody)
	if dbErr != nil {
		c.JSON(dbStatus, gin.H{"error": dbErr.Error()})
		return
	}

	c.JSON(dbStatus, gin.H{"data": dbUser})
}

func DeleteUser(c *gin.Context) {
	// Authenticate
	clerkUser, authenticatedErr := auth.AuthenticateClerkUser(c)
	if authenticatedErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": authenticatedErr.Error()})
		return
	}
	// Validate Request Body
	validatedRequestBody, validatedRequestErr := ValidateRequest(c)
	if validatedRequestErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validatedRequestErr.Error()})
		return
	}

	response, err := DeleteUserService(&clerkUser, validatedRequestBody)
	if err != nil {
		c.JSON(response, gin.H{"error": err.Error()})
		return
	}

	// Success
	c.JSON(response, gin.H{"message": "user deleted"})
}
