package decision

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/twjsanderson/decision_backend/internal/auth"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func ValidateRequest(c *gin.Context) (*models.Decision, error) {
	// Validate the JSON format
	var requestBody models.Decision
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		return nil, fmt.Errorf("invalid request body format")
	}

	// Validate required field(s)
	if requestBody.Id == "" {
		return nil, fmt.Errorf("missing or empty id field")
	}

	// Get the route path from the context
	routePath := c.FullPath()

	if routePath == "/decision/create/initial" {
		if requestBody.Title == "" ||
			requestBody.ChoiceType == "" ||
			requestBody.Problem == "" ||
			requestBody.IdealOutcome == "" ||
			requestBody.MaxCost == "" ||
			requestBody.RiskTolerance == "" ||
			requestBody.Timeline == "" {
			return nil, fmt.Errorf("missing or empty required field(s)")
		}
	}

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
	// Your logic for getting a decision
	c.JSON(http.StatusOK, gin.H{"message": "Decision fetched"})
}

// func UpdateDecision(c *gin.Context) {
// 	// Your logic for updating a decision
// 	c.JSON(http.StatusOK, gin.H{"message": "Decision updated"})
// }

// func DeleteDecision(c *gin.Context) {
// 	// Your logic for deleting a decision
// 	c.JSON(http.StatusOK, gin.H{"message": "Decision deleted"})
// }
