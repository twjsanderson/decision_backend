package routes

import (
    "github.com/gin-gonic/gin"

    "github.com/twjsanderson/decision_backend/pkg/user"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // User routes
    userRoutes := router.Group("/user")
    {
        userRoutes.POST("/create", user.CreateUser)
        userRoutes.GET("/get", user.GetUser)
        userRoutes.PUT("/update", user.UpdateUser)
        userRoutes.DELETE("/delete", user.DeleteUser)
    }

    // Decision routes
    // decisionRoutes := router.Group("/decision")
    // {
    //     decisionRoutes.POST("/create", handlers.CreateDecision)
    //     decisionRoutes.GET("/get", handlers.GetDecision)
    //     decisionRoutes.PUT("/update", handlers.UpdateDecision)
    //     decisionRoutes.DELETE("/delete", handlers.DeleteDecision)
    // }

    // Payment routes
    // paymentRoutes := router.Group("/payment")
    // {
    //     paymentRoutes.POST("/subscribe", handlers.CreateSubscription)
    //     paymentRoutes.POST("/cancel", handlers.CancelSubscription)
    //     paymentRoutes.GET("/status", handlers.GetSubscriptionStatus)
    // }

    return router
}
