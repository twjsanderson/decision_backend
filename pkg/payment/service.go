package payment

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/subscription"
	"log"
)

func CreateStripeSubscription() error {
	// Stripe API setup
	stripe.Key = "your-stripe-secret-key"

	// Sample data: Replace with actual user data
	customerID := "cus_example"

	// Create the subscription
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String("price_example"), // Replace with your price ID
			},
		},
	}

	sub, err := subscription.New(params)
	if err != nil {
		log.Println("Error creating subscription:", err)
		return err
	}

	log.Println("Subscription created:", sub.ID)
	return nil
}

func CancelStripeSubscription() error {
	// Cancel subscription logic
	// Example: you would cancel the subscription by ID here
	return nil
}
