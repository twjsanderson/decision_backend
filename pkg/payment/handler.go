package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSubscription(c *gin.Context) {
	// Call payment service to create a Stripe subscription
	// err := payment.CreateStripeSubscription() // Service function to handle Stripe logic
	// if err != nil {
	//     c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//     return
	// }
	c.JSON(http.StatusOK, gin.H{"message": "Subscription created"})
}

func CancelSubscription(c *gin.Context) {
	// Call payment service to cancel a Stripe subscription
	// err := payment.CancelStripeSubscription()
	// if err != nil {
	//     c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//     return
	// }
	c.JSON(http.StatusOK, gin.H{"message": "Subscription cancelled"})
}

func GetSubscriptionStatus(c *gin.Context) {
	// Call payment service to get status of user subscription
	// status (100: OK, 200: Error)
	// type: Free or Paid
	// paymentDate: future date or nil
	c.JSON(http.StatusOK, gin.H{
		"status":      100,
		"type":        "Free",
		"paymentDate": nil,
		"message":     "Here is your subscription status",
	})
}
