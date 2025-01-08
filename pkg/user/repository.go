package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/twjsanderson/decision_backend/internal/db"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func InsertUser(user *models.User) (int, error) {
	query := `INSERT INTO users (id, email, first_name, last_name) 
			  VALUES ($1, $2, $3, $4)`

	_, err := db.DB.Exec(context.Background(), query, user.Id, user.Email, user.FirstName, user.LastName)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to insert user: %v", err)
	}

	// Success
	return http.StatusCreated, nil
}

func GetUserById(id *string) (models.User, int, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id = $1"
	err := db.DB.QueryRow(context.Background(), query, *id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)

	// Check for the specific error message that represents no rows found
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, http.StatusNotFound, fmt.Errorf("user not found")
		}
		return user, http.StatusInternalServerError, fmt.Errorf("db error %v", err)
	}

	return user, http.StatusFound, nil
}

func DeleteUserById(id *string) (int, error) {
	var user models.User

	// Retrieve the user to ensure it exists before deletion
	queryGet := "SELECT * FROM users WHERE id = $1"
	err := db.DB.QueryRow(context.Background(), queryGet, *id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		fmt.Printf("njkpfhnfgd: %v", err)
		if err == pgx.ErrNoRows {
			return http.StatusNotFound, fmt.Errorf("user with id %s not found", *id)
		}
		return http.StatusInternalServerError, fmt.Errorf("failed to fetch user before deletion - %v", err)
	}

	// Delete the user from the database
	queryDelete := "DELETE FROM users WHERE id = $1"
	res, err := db.DB.Exec(context.Background(), queryDelete, *id)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to delete user with id %s - %v", *id, err)
	}
	fmt.Println("DeleteUserById after all: ", res, err)

	return http.StatusAccepted, nil
}
