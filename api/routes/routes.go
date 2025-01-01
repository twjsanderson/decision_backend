package routes

import (
    "github.com/gin-gonic/gin"
    "decision_backend/handlers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // User routes
    userRoutes := router.Group("/user")
    {
        userRoutes.POST("/create", handlers.CreateUser)
        userRoutes.GET("/get", handlers.GetUser)
        userRoutes.PUT("/update", handlers.UpdateUser)
        userRoutes.DELETE("/delete", handlers.DeleteUser)
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
