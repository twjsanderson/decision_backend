package decision

import (
	"fmt"
	"net/http"

	"github.com/twjsanderson/decision_backend/internal/models"
	"github.com/twjsanderson/decision_backend/pkg/permissions"
	"github.com/twjsanderson/decision_backend/pkg/user"
)

func ExtractDecision(completeDecision *models.Decision) models.DecisionResponse {
	return models.DecisionResponse{
		DecisionId: completeDecision.DecisionId,
		Title:      completeDecision.Title,
		Problem:    completeDecision.Problem,
		UserId:     completeDecision.Id,
	}
}

func CreateDecisionService(
	clerkUser *models.ClerkUser,
	requestBody *models.Decision,
) (models.DecisionResponse, int, error) {
	var onlyDecision models.DecisionResponse

	// Extract the embedded User from requestBody
	userReference := &requestBody.User

	authStatus, authErr := user.AuthorizeUserService(clerkUser, userReference, "CREATE")
	if authErr != nil {
		return onlyDecision, authStatus, authErr
	}
	// Get permissions for User
	if authStatus == http.StatusOK {
		userPermissions, permissionsStatus, permissionsErr := permissions.GetUserPermissionsById(&userReference.Id)
		if permissionsErr != nil {
			return onlyDecision, permissionsStatus, permissionsErr
		}

		// Check if user has enough points to create new decision
		if userPermissions.Max > 0 {
			decisionId, _, insertionErr := InsertDecision(requestBody)
			if insertionErr != nil {
				return onlyDecision, http.StatusInternalServerError, fmt.Errorf("failed to insert new decision - %v", insertionErr)
			}

			newDecision, newDecisionStatus, newDecisionErr := GetDecisionById(decisionId)
			if newDecisionErr != nil {
				return onlyDecision, newDecisionStatus, fmt.Errorf("failed to get new decision - %v", newDecisionErr)
			}

			// Update user_permissions with subtracted max value
			var max int = userPermissions.Max - 1
			permissionsStatus, permissionsErr := permissions.UpdateUserPermissions(max, userReference.Id)
			if permissionsErr != nil {
				return onlyDecision, permissionsStatus, permissionsErr
			}

			return ExtractDecision(&newDecision), http.StatusCreated, nil
		}
		if userPermissions.Max == 0 {
			return onlyDecision, http.StatusBadRequest, fmt.Errorf("user does not have enough credits to create new decisison")
		}
	}
	return onlyDecision, http.StatusInternalServerError, fmt.Errorf("an error occured while trying to create new decision for %v", requestBody.Id)
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

func GetDecisionService(
	clerkUser *models.ClerkUser,
	requestBody *models.Decision,
) (models.DecisionResponse, int, error) {
	var onlyDecision models.DecisionResponse

	// Extract the embedded User from requestBody
	userReference := &requestBody.User

	authStatus, authErr := user.AuthorizeUserService(clerkUser, userReference, "GET")
	if authErr != nil {
		return onlyDecision, authStatus, authErr
	}
	fmt.Print(&requestBody.DecisionId)
	decision, decisionStatus, decisionErr := GetDecisionById(requestBody.DecisionId)
	if decisionErr != nil {
		return onlyDecision, decisionStatus, fmt.Errorf("failed to get decision with id: %v, error: %v", requestBody.DecisionId, decisionErr)
	}

	return ExtractDecision(&decision), http.StatusFound, nil
}

func GetAllDecisionService(
	clerkUser *models.ClerkUser,
	requestBody *models.Decision,
) (int, error) {
	return 1, fmt.Errorf("GetAllDecisionService")
}
