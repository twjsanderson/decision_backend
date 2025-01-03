package user

import (
	"errors"

	"github.com/twjsanderson/decision_backend/internal/models"
)

func CreateUserService(user *models.User) (*models.User, error) {
	// Business logic to create the user (e.g., interacting with DB)
	// For example, you might want to validate the user data, hash a password, etc.

	if user.Id == "" {
		return user, errors.New("User id is required")
	}

	// Here you would typically interact with the database to save the user
	// For now, we'll just simulate a successful creation.

	// Simulate user being created
	// user.ID = someGeneratedID

	/// Returning nil indicates no error, meaning the user was successfully created.
	return user, nil
}

func GetUserService(id int) (models.User, error) {
	var user models.User

	return user, nil
}
