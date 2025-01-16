package decision

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/twjsanderson/decision_backend/internal/db"
	"github.com/twjsanderson/decision_backend/internal/models"
)

func InsertDecision(decision *models.Decision) (int, error) {
	query := `
		INSERT INTO decisions (user_id, title, choice_type, problem, justification, ideal_outcome, max_cost, risk_tolerance, timeline) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id` // Return the primary key (decision ID)

	var insertedDecisionID int
	err := db.DB.QueryRow(context.Background(), query,
		decision.Id, // This is the user_id (foreign key), inherited from ClerkUser
		decision.Title,
		decision.ChoiceType,
		decision.Problem,
		decision.Justification,
		decision.IdealOutcome,
		decision.MaxCost,
		decision.RiskTolerance,
		decision.Timeline,
	).Scan(&insertedDecisionID)

	if err != nil {
		return -1, fmt.Errorf("failed to insert decision: %v", err)
	}

	// Success
	return insertedDecisionID, nil
}

func GetDecisionById(id int) (models.Decision, error) {
	query := `
		SELECT 
			id, 
			user_id, 
			title, 
			choice_type, 
			problem, 
			justification, 
			ideal_outcome, 
			max_cost, 
			risk_tolerance, 
			timeline 
		FROM decisions 
		WHERE id = $1`

	var decision models.Decision

	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&decision.Id, // Maps to the `user_id` field (foreign key)
		&decision.User.ClerkUser.Id,
		&decision.Title,
		&decision.ChoiceType,
		&decision.Problem,
		&decision.Justification,
		&decision.IdealOutcome,
		&decision.MaxCost,
		&decision.RiskTolerance,
		&decision.Timeline,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return decision, fmt.Errorf("decision with id %d not found", id)
		}
		return decision, fmt.Errorf("failed to fetch decision: %v", err)
	}

	return decision, nil
}
