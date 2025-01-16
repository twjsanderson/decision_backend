package decision

import (
	"fmt"
	"net/http"

	"github.com/twjsanderson/decision_backend/internal/models"
	"github.com/twjsanderson/decision_backend/pkg/user"
)

func CreateDecisionService(
	clerkUser *models.ClerkUser,
	requestBody *models.Decision,
) (models.Decision, int, error) {
	var decision models.Decision

	// Extract the embedded User from requestBody
	userReference := &requestBody.User

	authStatus, authErr := user.AuthorizeUserService(clerkUser, userReference, "CREATE")
	if authErr != nil || authStatus != http.StatusOK {
		var decision models.Decision
		return decision, authStatus, authErr
	}

	decisionId, insertionErr := InsertDecision(requestBody)
	if insertionErr != nil && decisionId == -1 {
		return decision, http.StatusInternalServerError, fmt.Errorf("failed to insert new decision - %v", insertionErr)
	}

	newDecision, newDecisionErr := GetDecisionById(decisionId)
	if newDecisionErr != nil {
		return decision, http.StatusInternalServerError, fmt.Errorf("failed to get new decision - %v", newDecisionErr)
	}

	// code to build decision by chatGPT call and db insertion...

	return newDecision, http.StatusCreated, nil
}

func CompleteDecisionService(
	clerkUser *models.ClerkUser,
	requestBody *models.Decision,
) (int, error) {
	// // Extract the embedded User from Decision
	// userReference := &requestBody.User

	// authStatus, authErr := user.AuthorizeUserService(clerkUser, userReference, "CREATE")
	// if authErr != nil {
	// 	return authStatus, authErr
	// }

	// _, dbStatus, dbErr := user.GetUserById(&clerkUser.Id)
	// if dbErr != nil && dbStatus != http.StatusNotFound {
	// 	return dbStatus, fmt.Errorf("failed to fetch authenticated user from DB - %v", dbErr)
	// }
	return 1, fmt.Errorf("error")
	// code to build decision by chatGPT call and db insertion...
}
