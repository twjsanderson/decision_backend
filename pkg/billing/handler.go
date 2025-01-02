package handlers

import (
	"github.com/gin-gonic/gin"
)

func CreateBillingEntry(c *gin.Context) {
	// Insert new billing Entry to DB when payment service notifies of successful payment
	// Each entry would be created per month as per the subscription deal
}

func GetBillingHistory(c *gin.Context) {
	// Call DB for Users' billing Entries
}
