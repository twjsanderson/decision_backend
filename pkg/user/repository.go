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
	err := db.DB.QueryRow(context.Background(), query, *id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.IsAdmin)

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
	err := db.DB.QueryRow(context.Background(), queryGet, *id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.IsAdmin)
	if err != nil {
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

func UpdateUserData(user *models.User) (models.User, int, error) {
	var existingUser models.User

	// Check if the user exists
	queryGet := "SELECT * FROM users WHERE id = $1"
	err := db.DB.QueryRow(context.Background(), queryGet, user.Id).Scan(
		&existingUser.Id,
		&existingUser.Email,
		&existingUser.FirstName,
		&existingUser.LastName,
		&existingUser.IsAdmin,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.User{}, http.StatusNotFound, fmt.Errorf("user with id %s not found", user.Id)
		}
		return models.User{}, http.StatusInternalServerError, fmt.Errorf("failed to fetch user before update - %v", err)
	}

	// Update only the fields that are not nil
	queryUpdate := `
		UPDATE users 
		SET email = COALESCE($1, email),
		    first_name = COALESCE($2, first_name),
		    last_name = COALESCE($3, last_name), 
			is_admin = COALESCE($4, is_admin)
		WHERE id = $5
	`
	_, err = db.DB.Exec(context.Background(), queryUpdate, user.Email, user.FirstName, user.LastName, user.IsAdmin, user.Id)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, fmt.Errorf("failed to update user with id %s - %v", user.Id, err)
	}

	// Fetch the updated user from the database
	var updatedUser models.User
	err = db.DB.QueryRow(context.Background(), queryGet, user.Id).Scan(
		&updatedUser.Id,
		&updatedUser.Email,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.IsAdmin,
	)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, fmt.Errorf("failed to fetch updated user with id %s - %v", user.Id, err)
	}

	// Success: Return the updated user (not a pointer)
	return updatedUser, http.StatusOK, nil
}
