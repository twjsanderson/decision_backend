package decision

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/twjsanderson/decision_backend/internal/auth"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func ValidateRequest(c *gin.Context) (*models.Decision, error) {
	var requestBody models.Decision

	// Get the route path from the context
	routePath := c.FullPath()

	// get user_id (as id) from params
	userId := c.Query("id")
	if userId == "" {
		return nil, fmt.Errorf("missing or empty required field(s)")
	}
	requestBody.Id = userId // assign userId from params to requestBody

	// Validate the JSON format for all req. except GET
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		return nil, fmt.Errorf("invalid request body format")
	}

	if routePath == "/decision/create" {
		if requestBody.Id == "" ||
			requestBody.Title == "" ||
			requestBody.Problem == "" {
			return nil, fmt.Errorf("missing or empty required field(s)")
		}
	}

	// handle /decision/update body

	// Return the valid request
	return &requestBody, nil
}

func CreateDecision(c *gin.Context) {
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

	decision, status, err := CreateDecisionService(&clerkUser, validatedRequestBody)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	// Success
	c.JSON(status, gin.H{"data": decision, "message": "initial decision created"})
}

func CompleteDecision(c *gin.Context) {
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

	status, err := CompleteDecisionService(&clerkUser, validatedRequestBody)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	// Success
	c.JSON(status, gin.H{"message": "opinions created"})
}

func GetDecision(c *gin.Context) {
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

	dbDecision, dbStatus, dbErr := GetDecisionService(&clerkUser, validatedRequestBody)
	if dbErr != nil {
		c.JSON(dbStatus, gin.H{"error": dbErr.Error()})
		return
	}

	c.JSON(dbStatus, gin.H{"data": dbDecision})
}

func UpdateDecision(c *gin.Context) {
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
	updatedDecision, updatedDecisionStatus, err := UpdateDecisionService(&clerkUser, validatedRequestBody)
	if err != nil {
		c.JSON(updatedDecisionStatus, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedDecision})
}

func DeleteDecision(c *gin.Context) {
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
	status, err := DeleteDecisionService(&clerkUser, validatedRequestBody)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	// Success
	c.JSON(status, gin.H{"message": "decision deleted"})
}
