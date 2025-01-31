package decision

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/twjsanderson/decision_backend/internal/db"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func InsertDecision(decision *models.Decision) (int, int, error) {
	query := `
		INSERT INTO decisions (user_id, title, problem) 
		VALUES ($1, $2, $3)
		RETURNING id` // Return the primary key (decision ID)

	var insertedDecisionID int
	err := db.DB.QueryRow(context.Background(), query,
		decision.Id, // This is the user_id (foreign key), inherited from ClerkUser
		decision.Title,
		decision.Problem,
		// decision.Justification,
		// decision.IdealOutcome,
		// decision.MaxCost,
		// decision.RiskTolerance,
		// decision.Timeline,
	).Scan(&insertedDecisionID)

	if err != nil {
		return -1, http.StatusInternalServerError, fmt.Errorf("failed to insert decision: %v", err)
	}

	// Success
	return insertedDecisionID, http.StatusNotFound, nil
}

func GetDecisionById(id int) (models.Decision, int, error) {
	query := `
		SELECT 
			id, 
			user_id, 
			title, 
			problem
		FROM decisions 
		WHERE id = $1`

	var decision models.Decision

	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&decision.DecisionId,
		&decision.User.ClerkUser.Id,
		&decision.Title,
		&decision.Problem,
		// &decision.Justification,
		// &decision.IdealOutcome,
		// &decision.MaxCost,
		// &decision.RiskTolerance,
		// &decision.Timeline,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return decision, http.StatusNotFound, fmt.Errorf("decision with id %d not found", id)
		}
		return decision, http.StatusInternalServerError, fmt.Errorf("failed to fetch decision: %v", err)
	}

	return decision, http.StatusFound, nil
}
