package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/twjsanderson/decision_backend/pkg/decision"
	"github.com/twjsanderson/decision_backend/pkg/user"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Add CORS middleware before defining routes
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // The frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// User routes
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/create", user.CreateUser)
		userRoutes.GET("/get", user.GetUser)
		userRoutes.PUT("/update", user.UpdateUser)
		userRoutes.DELETE("/delete", user.DeleteUser)
	}

	// Decision routes
	decisionRoutes := router.Group("/decision")
	{
		decisionRoutes.POST("/create", decision.CreateDecision)
		decisionRoutes.POST("/complete", decision.CompleteDecision)
		decisionRoutes.GET("/get", decision.GetDecision)
		decisionRoutes.PUT("/update", decision.UpdateDecision)
		// decisionRoutes.DELETE("/delete", handlers.DeleteDecision)
	}

	// Payment routes
	// paymentRoutes := router.Group("/payment")
	// {
	//     paymentRoutes.POST("/subscribe", handlers.CreateSubscription)
	//     paymentRoutes.POST("/cancel", handlers.CancelSubscription)
	//     paymentRoutes.GET("/status", handlers.GetSubscriptionStatus)
	// }

	return router
}
