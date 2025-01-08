package user

import (
	"fmt"
	"net/http"

	"github.com/twjsanderson/decision_backend/internal/auth"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func AuthorizeUserService(
	clerkUser *models.ClerkUser,
	requestBody *models.User,
	operation string,
) (int, error) {
	// Fetch user from DB
	dbUser, httpStatus, dbErr := GetUserById(&clerkUser.Id)
	if dbErr != nil && operation != "CREATE" {
		fmt.Print(operation, "CREATE")
		return httpStatus, fmt.Errorf("failed to fetch authenticated user from DB - %v", dbErr)
	}

	// Check authorization
	authorized := auth.AuthorizeUserOperation(clerkUser, &dbUser, requestBody, operation)
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
	if dbStatus == http.StatusNotFound {
		insertionStatus, insertionErr := InsertUser(requestBody)
		if insertionErr != nil {
			return insertionStatus, insertionErr
		}
		return insertionStatus, nil
	}
	return dbStatus, fmt.Errorf("user already exists")
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
	return dbUser, httpStatus, dbErr
}

func DeleteUserService(
	clerkUser *models.ClerkUser,
	requestBody *models.User,
) (int, error) {
	response, err := AuthorizeUserService(clerkUser, requestBody, "DELETE")
	if err != nil {
		return response, err
	}
	_, deletionErr := DeleteUserById(&requestBody.Id)
	if deletionErr != nil {
		return response, deletionErr
	}
	return response, nil
}
