package user

import (
	"fmt"
	"net/http"

	"github.com/twjsanderson/decision_backend/internal/models"
	"github.com/twjsanderson/decision_backend/pkg/permissions"
)

func AuthorizeUserOperation(
	clerkUser *models.ClerkUser,
	dbUser *models.User,
	requestBody *models.User,
	operation string,
) bool {
	if operation == "CREATE" || operation == "DELETE" {
		if clerkUser.Id == requestBody.Id {
			return true
		}
	}
	if operation == "GET" {
		if dbUser.IsAdmin || dbUser.Id == requestBody.Id {
			return true
		}
	}
	if operation == "UPDATE" {
		if dbUser.IsAdmin || clerkUser.Id == requestBody.Id { // should just be IsAdmin in Prod
			return true
		}
	}
	return false
}

func AuthorizeUserService(
	clerkUser *models.ClerkUser,
	requestBody *models.User,
	operation string,
) (int, error) {
	// Fetch user from DB
	dbUser, httpStatus, dbErr := GetUserById(&clerkUser.Id)
	if dbErr != nil && operation != "CREATE" {
		return httpStatus, fmt.Errorf("failed to fetch authenticated user from DB - %v", dbErr)
	}
	// Check authorization
	authorized := AuthorizeUserOperation(clerkUser, &dbUser, requestBody, operation)
	if !authorized {
		return http.StatusUnauthorized, fmt.Errorf("user is not authorized for %v operation", operation)
	}
	// Success
	return http.StatusOK, nil
}

func CreateUserService(
	clerkUser *models.ClerkUser,
	requestBody *models.User,
) (int, error) {
	authStatus, authErr := AuthorizeUserService(clerkUser, requestBody, "CREATE")
	if authErr != nil {
		return authStatus, authErr
	}

	_, dbStatus, dbErr := GetUserById(&clerkUser.Id)
	if dbErr != nil && dbStatus != http.StatusNotFound {
		return dbStatus, fmt.Errorf("failed to fetch authenticated user from DB - %v", dbErr)
	}

	if dbStatus != http.StatusNotFound {
		return http.StatusConflict, fmt.Errorf("user already exists")
	}

	// Ensure IsAdmin is false for the new user
	requestBody.IsAdmin = false

	// Insert the new user
	insertionStatus, insertionErr := InsertUser(requestBody)
	if insertionErr != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to insert user - %v", insertionErr)
	}

	var defaultPermissions models.UserPermissions = models.UserPermissions{
		UserId: clerkUser.Id,
		Max:    3,
	}

	// Insert default permissions for new user
	_, permissionsStatus, permissionsErr := permissions.InsertUserPermissions(defaultPermissions)
	if permissionsErr != nil {
		return permissionsStatus, permissionsErr
	}

	// Success
	return insertionStatus, nil
}

func GetUserService(
	clerkUser *models.ClerkUser,
	requestBody *models.User,
) (models.User, int, error) {
	var user models.User
	authStatus, authErr := AuthorizeUserService(clerkUser, requestBody, "GET")
	if authErr != nil {
		return user, authStatus, authErr
	}
	// Fetch user from DB
	dbUser, httpStatus, dbErr := GetUserById(&clerkUser.Id)
	if dbErr != nil && httpStatus != http.StatusNotFound {
		return user, httpStatus, fmt.Errorf("failed to fetch authenticated user from DB - %v", dbErr)
	}
	return dbUser, http.StatusOK, dbErr
}

func DeleteUserService(
	clerkUser *models.ClerkUser,
	requestBody *models.User,
) (int, error) {
	response, err := AuthorizeUserService(clerkUser, requestBody, "DELETE")
	if err != nil {
		return response, err
	}
	_, permissionsDeletionErr := permissions.DeletePermissionsById(&requestBody.Id)
	if permissionsDeletionErr != nil {
		return response, permissionsDeletionErr
	}
	_, deletionErr := DeleteUserById(&requestBody.Id)
	if deletionErr != nil {
		return response, deletionErr
	}
	return response, nil
}

func UpdateUserService(
	clerkUser *models.ClerkUser,
	requestBody *models.User,
) (models.User, int, error) {
	var user models.User

	authStatus, authErr := AuthorizeUserService(clerkUser, requestBody, "UPDATE")
	if authErr != nil {
		return user, authStatus, authErr
	}
	// Update user
	updatedUser, updateStatus, updateErr := UpdateUserData(requestBody)
	if updateErr != nil {
		return user, updateStatus, fmt.Errorf("failed to insert user - %v", updateErr)
	}
	return updatedUser, updateStatus, updateErr
}
