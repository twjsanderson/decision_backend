package permissions

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/twjsanderson/decision_backend/internal/db"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func UpdateUserPermissions(max int, pkg string, lastChecked time.Time, user_id string) (int, error) {
	queryUpdate := `
		UPDATE user_permissions 
		SET max = COALESCE($1, max),
			package = COALESCE($2, package),
			last_checked = COALESCE($3, last_checked)
		WHERE user_id = $4
	`

	_, err := db.DB.Exec(context.Background(), queryUpdate, max, pkg, lastChecked, user_id) // Make sure user_id is passed as a string, not an integer
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("UpdateUserPermissions error - failed to update user_permissions with user_id %s - %v", user_id, err)
	}
	return http.StatusAccepted, nil
}

func GetUserPermissionsById(user_id *string) (models.UserPermissions, int, error) {
	var permissions models.UserPermissions
	var packageNull sql.NullString
	var lastCheckedNull sql.NullTime

	query := "SELECT id, user_id, max, package, last_checked FROM user_permissions WHERE user_id = $1"
	err := db.DB.QueryRow(context.Background(), query, *user_id).Scan(
		&permissions.Id,
		&permissions.UserId,
		&permissions.Max,
		&permissions.Package,
		&permissions.LastChecked,
	)

	// Convert NULL values to appropriate types
	if packageNull.Valid {
		permissions.Package = packageNull.String
	} else {
		permissions.Package = "" // Default empty string if NULL
	}

	if lastCheckedNull.Valid {
		permissions.LastChecked = lastCheckedNull.Time
	} else {
		permissions.LastChecked = time.Time{} // Default zero time if NULL
	}

	if err != nil {
		fmt.Print(permissions)
		if err == pgx.ErrNoRows {
			return permissions, http.StatusNotFound, fmt.Errorf("GetUserPermissionsById error - permissions not found")
		}
		return permissions, http.StatusInternalServerError, fmt.Errorf("GetUserPermissionsById error - %v", err)
	}

	// Success
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
		return models.UserPermissions{}, http.StatusInternalServerError, fmt.Errorf("InsertUserPermissions error - failed to insert user permissions: %v", err)
	}

	// Success
	return insertedPermissions, http.StatusCreated, nil
}

func DeletePermissionsById(id *string) (int, error) {
	queryDelete := "DELETE FROM user_permissions WHERE user_id = $1"
	res, err := db.DB.Exec(context.Background(), queryDelete, *id)
	fmt.Println("DeletePermissionsById after all: ", res, err)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("DeletePermissionsById error - failed to delete user with id %s - %v", *id, err)
	}

	return http.StatusAccepted, nil
}
