package user

import (
	"context"
	"fmt"

	"github.com/twjsanderson/decision_backend/internal/db"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func InsertUser(user models.User) error {
	query := `INSERT INTO users (id, email, first_name, last_name) 
			  VALUES ($1, $2, $3, $4)`

	_, err := db.DB.Exec(context.Background(), query, user.Id, user.Email, user.FirstName, user.LastName)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	// Success
	return nil
}

func GetUserById(id string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id = $1"
	err := db.DB.QueryRow(context.Background(), query, id)
	if err != nil {
		return user, fmt.Errorf("failed to fetch user: %v", err)
	}

	// Return the user and no error if the query was successful
	return user, nil
}

func UpdateUser(user models.User) (models.User, error) {

}
