package permissions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/twjsanderson/decision_backend/internal/db"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func UpdateUserPermissions(max int, user_id string) (int, error) {
	queryUpdate := `
		UPDATE user_permissions 
		SET max = $1
		WHERE user_id = $2
	`

	_, err := db.DB.Exec(context.Background(), queryUpdate, max, user_id) // Make sure user_id is passed as a string, not an integer
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to update user_permissions with user_id %s - %v", user_id, err)
	}
	return http.StatusAccepted, nil
}

func GetUserPermissionsById(user_id *string) (models.UserPermissions, int, error) {
	var permissions models.UserPermissions

	query := "SELECT id, user_id, max FROM user_permissions WHERE user_id = $1"
	err := db.DB.QueryRow(context.Background(), query, *user_id).Scan(
		&permissions.Id,
		&permissions.UserId,
		&permissions.Max,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return permissions, http.StatusNotFound, fmt.Errorf("permissions not found")
		}
		return permissions, http.StatusInternalServerError, fmt.Errorf("db error %v", err)
	}

	return permissions, http.StatusOK, nil
}

func InsertUserPermissions(permissions models.UserPermissions) (models.UserPermissions, int, error) {
	query := `INSERT INTO user_permissions (user_id, max) 
			  VALUES ($1, $2) RETURNING id, user_id, max`

	// Execute the query and fetch the inserted record
	var insertedPermissions models.UserPermissions
	err := db.DB.QueryRow(context.Background(), query, permissions.UserId, permissions.Max).
		Scan(&insertedPermissions.Id, &insertedPermissions.UserId, &insertedPermissions.Max)

	if err != nil {
		return models.UserPermissions{}, http.StatusInternalServerError, fmt.Errorf("failed to insert user permissions: %v", err)
	}

	// Success
	return insertedPermissions, http.StatusCreated, nil
}

func DeletePermissionsById(id *string) (int, error) {
	queryDelete := "DELETE FROM user_permissions WHERE user_id = $1"
	res, err := db.DB.Exec(context.Background(), queryDelete, *id)
	fmt.Println("DeletePermissionsById after all: ", res, err)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to delete user with id %s - %v", *id, err)
	}

	return http.StatusAccepted, nil
}
