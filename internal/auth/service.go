package auth

import (
	"fmt"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	clerkUser "github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"

	"github.com/twjsanderson/decision_backend/internal/config"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func AuthenticateClerkUser(c *gin.Context) (models.ClerkUser, error) {
	appConfig := config.LoadConfig()
	clerk.SetKey(appConfig.CLERK_API_KEY)

	ctx := c.Request.Context()
	authHeader := c.GetHeader("Authorization")

	var emptyUser models.ClerkUser

	// Check if the Authorization header is provided
	if authHeader == "" {
		return emptyUser, fmt.Errorf("missing Authorization header")
	}

	// Extract the Bearer token from the Authorization header
	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return emptyUser, fmt.Errorf("invalid Authorization header format")
	}
	bearer := authHeader[len(bearerPrefix):]

	// Verify the jwt
	claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
		Token: bearer,
	})

	if err != nil {
		return emptyUser, fmt.Errorf("failed to validate token")
	}

	// Retrieve user details from Clerk using the subject (user ID) in claims
	userDetails, err := clerkUser.Get(ctx, claims.Subject)
	if err != nil {
		return emptyUser, fmt.Errorf("failed to fetch user details: %w", err)
	}

	email := ""
	if len(userDetails.EmailAddresses) > 0 {
		email = userDetails.EmailAddresses[0].EmailAddress
	}

	return models.ClerkUser{
			Id:        userDetails.ID,
			Email:     email,
			FirstName: userDetails.FirstName,
			LastName:  userDetails.LastName},
		nil
}

func AuthorizeUserOperation(
	clerkUser *models.ClerkUser,
	dbUser *models.User,
	requestBody *models.User,
	operation string,
) bool {
	if operation == "CREATE" || operation == "DELETE" {
		if clerkUser.Id == requestBody.Id && clerkUser.Email == requestBody.Email {
			return true
		}
	}
	if operation == "GET" {
		if dbUser.IsAdmin || dbUser.Id == requestBody.Id {
			return true
		}
	}
	if operation == "UPDATE" {
		if dbUser.IsAdmin {
			// specific fields only!
			return true
		}
		if dbUser.Id == requestBody.Id {
			// specific fields only!
			return true
		}
	}
	return false
}
