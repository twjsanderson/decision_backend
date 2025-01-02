package middleware

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/gin-gonic/gin"
    
	"github.com/twjsanderson/decision_backend/internal/config"
)

func AuthMiddleware() gin.HandlerFunc {
	appConfig := config.LoadConfig()
	clerk.SetKey(appConfig.CLERK_API_KEY)

	return func(c *gin.Context) {
		// Clerk middleware to validate the Authorization header
		handler := clerkhttp.WithHeaderAuthorization()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		}))

		// Use the middleware
		handler.ServeHTTP(c.Writer, c.Request)

		// Validate the session and retrieve claims
		claims, ok := clerk.SessionClaimsFromContext(c.Request.Context())
		if !ok {
			// Return unauthorized response if authentication fails
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // Stop processing further
			return
		}

		// Add claims to the context for downstream handlers
		c.Set("user", claims)
		c.Next() // Proceed to the next handler
	}
}
