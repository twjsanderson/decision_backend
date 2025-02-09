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
	).Scan(&insertedDecisionID)

	//.. add more isnertions

	if err != nil {
		return -1, http.StatusInternalServerError, fmt.Errorf("failed to insert decision: %v", err)
	}

	// Success
	return insertedDecisionID, http.StatusNotFound, nil
}

func GetDecisionById(id int) (models.Decision, int, error) {
	var decision models.Decision

	decision.ParetoAnalysis = &models.ParetoAnalysis{}
	decision.SwotAnalysis = &models.SwotAnalysis{}
	decision.BayesianDecision = &models.BayesianDecision{}
	decision.DecisionTree = &models.DecisionTree{}
	decision.AHP = &models.AnalyticHierarchyProcess{}
	decision.FirstPrinciples = &models.FirstPrinciples{}
	decision.FuzzyLogic = &models.FuzzyLogicDecisionMaking{}
	decision.CostBenefitAnalysis = &models.CostBenefitAnalysis{}

	baseQuery := `
		SELECT 
			id, 
			user_id, 
			title, 
			problem
		FROM decisions 
		WHERE id = $1`

	err := db.DB.QueryRow(context.Background(), baseQuery, id).Scan(
		&decision.DecisionId,
		&decision.User.ClerkUser.Id,
		&decision.Title,
		&decision.Problem,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return decision, http.StatusNotFound, fmt.Errorf("decision with id %d not found", id)
		}
		return decision, http.StatusInternalServerError, fmt.Errorf("failed to fetch decision with baseQuery: %v", err)
	}

	paretoQuery := `
		SELECT  
			possible_causes, 
			impact_measures,
			expected_outcome
		FROM pareto_analysis 
		WHERE decision_id = $1`

	paretoErr := db.DB.QueryRow(context.Background(), paretoQuery, id).Scan(
		&decision.ParetoAnalysis.PossibleCauses,
		&decision.ParetoAnalysis.ImpactMeasures,
		&decision.ParetoAnalysis.ExpectedOutcome,
	)
	fmt.Printf("Pareto Analysis: %+v\n", decision.ParetoAnalysis)
	if paretoErr != nil {
		if paretoErr == pgx.ErrNoRows {
			// If no Pareto analysis exists, return empty fields instead of nil
			decision.ParetoAnalysis.PossibleCauses = []string{}
			decision.ParetoAnalysis.ImpactMeasures = []string{}
			decision.ParetoAnalysis.ExpectedOutcome = ""
		} else {
			return decision, http.StatusInternalServerError, fmt.Errorf("failed to fetch decision with paretoQuery: %v", paretoErr)
		}
	}

	swotQuery := `
		SELECT  
			strengths, 
			weaknesses,
			opportunities,
			threats
		FROM swot_analysis 
		WHERE decision_id = $1`

	swotErr := db.DB.QueryRow(context.Background(), swotQuery, id).Scan(
		&decision.SwotAnalysis.Strengths,
		&decision.SwotAnalysis.Weaknesses,
		&decision.SwotAnalysis.Opportunities,
		&decision.SwotAnalysis.Threats,
	)
	fmt.Printf("SWOT: %+v\n", decision.SwotAnalysis)
	if swotErr != nil {
		if swotErr == pgx.ErrNoRows {
			// If no SWOT analysis exists, return empty fields instead of nil
			decision.SwotAnalysis.Strengths = []string{}
			decision.SwotAnalysis.Weaknesses = []string{}
			decision.SwotAnalysis.Opportunities = []string{}
			decision.SwotAnalysis.Threats = []string{}
		} else {
			return decision, http.StatusInternalServerError, fmt.Errorf("failed to fetch decision with swotQuery: %v", swotErr)
		}
	}

	bayesQuery := `
		SELECT  
			hypothesis, 
			prior_probability,
			evidence,
			likelihood_hypothesis_true,
			likelihood_hypothesis_false,
			expected_outcome
		FROM bayesian_decision 
		WHERE decision_id = $1`

	bayesErr := db.DB.QueryRow(context.Background(), bayesQuery, id).Scan(
		&decision.BayesianDecision.Hypothesis,
		&decision.BayesianDecision.PriorProbability,
		&decision.BayesianDecision.Evidence,
		&decision.BayesianDecision.LikelihoodHypothesisTrue,
		&decision.BayesianDecision.LikelihoodHypothesisFalse,
		&decision.BayesianDecision.ExpectedOutcome,
	)
	fmt.Printf("BAYES: %+v\n", decision.BayesianDecision)
	if bayesErr != nil {
		if bayesErr == pgx.ErrNoRows {
			// If no Bayes exists, return empty fields instead of nil
			decision.BayesianDecision.Hypothesis = ""
			decision.BayesianDecision.PriorProbability = 0.0
			decision.BayesianDecision.Evidence = ""
			decision.BayesianDecision.LikelihoodHypothesisTrue = 0.0
			decision.BayesianDecision.LikelihoodHypothesisFalse = 0.0
			decision.BayesianDecision.ExpectedOutcome = ""
		} else {
			return decision, http.StatusInternalServerError, fmt.Errorf("failed to fetch decision with bayesQuery: %v", bayesErr)
		}
	}

	decisionTree := `
		SELECT  
			options
		FROM decision_tree 
		WHERE decision_id = $1`

	decisitonTreeErr := db.DB.QueryRow(context.Background(), decisionTree, id).Scan(
		&decision.DecisionTree.Options,
	)
	fmt.Printf("Decision TREE: %+v\n", decision.DecisionTree)
	if decisitonTreeErr != nil {
		if decisitonTreeErr == pgx.ErrNoRows {
			// If no Decision Tree exists, return empty fields instead of nil
			decision.DecisionTree.Options = []byte("{}")
		} else {
			return decision, http.StatusInternalServerError, fmt.Errorf("failed to fetch decision with decisionTree: %v", decisitonTreeErr)
		}
	}

	analyticHierarchy := `
		SELECT  
			criteria,
			alternatives
		FROM analytic_hierarchy_process 
		WHERE decision_id = $1`

	analyticHierarchyErr := db.DB.QueryRow(context.Background(), analyticHierarchy, id).Scan(
		&decision.AHP.Criteria,
		&decision.AHP.Alternatives,
	)
	fmt.Printf("analyticHierarchy: %+v\n", decision.AHP)
	if analyticHierarchyErr != nil {
		if analyticHierarchyErr == pgx.ErrNoRows {
			// If no analyticHierarchy exists, return empty fields instead of nil
			decision.AHP.Criteria = []byte("{}")
			decision.AHP.Alternatives = []byte("{}")
		} else {
			return decision, http.StatusInternalServerError, fmt.Errorf("failed to fetch decision with analyticHierarchy: %v", analyticHierarchyErr)
		}
	}

	firstPrinciplesQuery := `
		SELECT  
			criteria,
			alternatives
		FROM analytic_hierarchy_process 
		WHERE decision_id = $1`

	firstPrinciplesErr := db.DB.QueryRow(context.Background(), firstPrinciplesQuery, id).Scan(
		&decision.FirstPrinciples.Assumptions,
		&decision.FirstPrinciples.FundamentalFacts,
		&decision.FirstPrinciples.ReconstructedSolution,
	)
	fmt.Printf("firstPrinciples: %+v\n", decision.FirstPrinciples)
	if firstPrinciplesErr != nil {
		if firstPrinciplesErr == pgx.ErrNoRows {
			// If no firstPrinciples exists, return empty fields instead of nil
			decision.FirstPrinciples.Assumptions = []byte("{}")
			decision.FirstPrinciples.FundamentalFacts = []byte("{}")
			decision.FirstPrinciples.ReconstructedSolution = ""
		} else {
			return decision, http.StatusInternalServerError, fmt.Errorf("failed to fetch decision with firstPrinciplesQuery: %v", firstPrinciplesErr)
		}
	}

	fuzzyQuery := `
		SELECT  
			fuzzy_variables,
			decision_threshold
		FROM fuzzy_logic 
		WHERE decision_id = $1
	`

	fuzzyErr := db.DB.QueryRow(context.Background(), fuzzyQuery, id).Scan(
		&decision.FuzzyLogic.FuzzyVariables,
		&decision.FuzzyLogic.DecisionThreshold,
	)
	fmt.Printf("fuzzyQuery: %+v\n", decision.FuzzyLogic)
	if fuzzyErr != nil {
		if fuzzyErr == pgx.ErrNoRows {
			// If no FuzzyLogic exists, return empty fields instead of nil
			decision.FuzzyLogic.FuzzyVariables = []byte("{}")
			decision.FuzzyLogic.DecisionThreshold = 0.0
		} else {
			return decision, http.StatusInternalServerError, fmt.Errorf("failed to fetch decision with fuzzyQuery: %v", fuzzyErr)
		}
	}

	costBenefitQuery := `
		SELECT  
			fuzzy_variables,
			decision_threshold
		FROM fuzzy_logic 
		WHERE decision_id = $1
	`

	costBenefitErr := db.DB.QueryRow(context.Background(), costBenefitQuery, id).Scan(
		&decision.CostBenefitAnalysis.Costs,
		&decision.CostBenefitAnalysis.Benefits,
		&decision.CostBenefitAnalysis.DiscountRate,
		&decision.CostBenefitAnalysis.TimeHorizon,
	)
	fmt.Printf("costBenefitQuery: %+v\n", decision.CostBenefitAnalysis)
	if costBenefitErr != nil {
		if costBenefitErr == pgx.ErrNoRows {
			// If no CostBenefitAnalysis exists, return empty fields instead of nil
			decision.CostBenefitAnalysis.Costs = []byte("{}")
			decision.CostBenefitAnalysis.Benefits = []byte("{}")
			decision.CostBenefitAnalysis.DiscountRate = 0.0
			decision.CostBenefitAnalysis.TimeHorizon = 0
		} else {
			return decision, http.StatusInternalServerError, fmt.Errorf("failed to fetch decision with costBenefitQuery: %v", costBenefitErr)
		}
	}

	return decision, http.StatusFound, nil
}

func DeleteDecisionById(id int) (int, error) {
	query := `
		DELETE FROM decisions
		WHERE id = $1
	`
	_, err := db.DB.Exec(context.Background(), query, id)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to fetch decision: %v", err)
	}

	return http.StatusFound, nil
}

func UpdateExistingDecision(decision *models.Decision) (int, int, error) {
	decisionID := decision.DecisionId

	tx, err := db.DB.Begin(context.Background()) // Start transaction
	if err != nil {
		return decisionID, http.StatusInternalServerError, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(context.Background()) // Ensure rollback on failure

	// Update Pareto Analysis
	if decision.ParetoAnalysis != nil {
		query := `
			UPDATE pareto_analysis
			SET possible_causes = $1, impact_measures = $2, expected_outcome = $3
			WHERE decision_id = $4
		`
		_, err := tx.Exec(context.Background(), query, decision.ParetoAnalysis.PossibleCauses, decision.ParetoAnalysis.ImpactMeasures, decision.ParetoAnalysis.ExpectedOutcome, decisionID)
		if err != nil {
			return decisionID, http.StatusInternalServerError, fmt.Errorf("failed to update pareto_analysis: %v", err)
		}
	}

	// Update SWOT Analysis
	if decision.SwotAnalysis != nil {
		query := `
			UPDATE swot_analysis
			SET strengths = $1, weaknesses = $2, opportunities = $3, threats = $4
			WHERE decision_id = $5
		`
		_, err := tx.Exec(context.Background(), query, decision.SwotAnalysis.Strengths, decision.SwotAnalysis.Weaknesses, decision.SwotAnalysis.Opportunities, decision.SwotAnalysis.Threats, decisionID)
		if err != nil {
			return decisionID, http.StatusInternalServerError, fmt.Errorf("failed to update swot_analysis: %v", err)
		}
	}

	// Update Bayesian Decision Theory
	if decision.BayesianDecision != nil {
		query := `
			UPDATE bayesian_decision
			SET hypothesis = $1, prior_probability = $2, evidence = $3, 
				likelihood_hypothesis_true = $4, likelihood_hypothesis_false = $5, expected_outcome = $6
			WHERE decision_id = $7
		`
		_, err := tx.Exec(context.Background(), query, decision.BayesianDecision.Hypothesis, decision.BayesianDecision.PriorProbability, decision.BayesianDecision.Evidence,
			decision.BayesianDecision.LikelihoodHypothesisTrue, decision.BayesianDecision.LikelihoodHypothesisFalse, decision.BayesianDecision.ExpectedOutcome, decisionID)
		if err != nil {
			return decisionID, http.StatusInternalServerError, fmt.Errorf("failed to update bayesian_decision: %v", err)
		}
	}

	// Update Decision Trees
	if decision.DecisionTree != nil {
		query := `
			UPDATE decision_tree
			SET options = $1
			WHERE decision_id = $2
		`
		_, err := tx.Exec(context.Background(), query, decision.DecisionTree.Options, decisionID)
		if err != nil {
			return decisionID, http.StatusInternalServerError, fmt.Errorf("failed to update decision_tree: %v", err)
		}
	}

	// Update Analytic Hierarchy Process
	if decision.AHP != nil {
		query := `
			UPDATE analytic_hierarchy_process
			SET criteria = $1, alternatives = $2
			WHERE decision_id = $3
		`
		_, err := tx.Exec(context.Background(), query, decision.AHP.Criteria, decision.AHP.Alternatives, decisionID)
		if err != nil {
			return decisionID, http.StatusInternalServerError, fmt.Errorf("failed to update analytic_hierarchy_process: %v", err)
		}
	}

	// Update First Principles Thinking
	if decision.FirstPrinciples != nil {
		query := `
			UPDATE first_principles
			SET assumptions = $1, fundamental_facts = $2, reconstructed_solution = $3
			WHERE decision_id = $4
		`
		_, err := tx.Exec(context.Background(), query, decision.FirstPrinciples.Assumptions, decision.FirstPrinciples.FundamentalFacts, decision.FirstPrinciples.ReconstructedSolution, decisionID)
		if err != nil {
			return decisionID, http.StatusInternalServerError, fmt.Errorf("failed to update first_principles: %v", err)
		}
	}

	// Update Fuzzy Logic Decision Making
	if decision.FuzzyLogic != nil {
		query := `
			UPDATE fuzzy_logic
			SET fuzzy_variables = $1, decision_threshold = $2
			WHERE decision_id = $3
		`
		_, err := tx.Exec(context.Background(), query, decision.FuzzyLogic.FuzzyVariables, decision.FuzzyLogic.DecisionThreshold, decisionID)
		if err != nil {
			return decisionID, http.StatusInternalServerError, fmt.Errorf("failed to update fuzzy_logic: %v", err)
		}
	}

	// Update Cost-Benefit Analysis
	if decision.CostBenefitAnalysis != nil {
		query := `
			UPDATE cost_benefit_analysis
			SET costs = $1, benefits = $2, discount_rate = $3, time_horizon = $4
			WHERE decision_id = $5
		`
		_, err := tx.Exec(context.Background(), query, decision.CostBenefitAnalysis.Costs, decision.CostBenefitAnalysis.Benefits, decision.CostBenefitAnalysis.DiscountRate, decision.CostBenefitAnalysis.TimeHorizon, decisionID)
		if err != nil {
			return decisionID, http.StatusInternalServerError, fmt.Errorf("failed to update cost_benefit_analysis: %v", err)
		}
	}

	// If everything succeeds, commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return decisionID, http.StatusInternalServerError, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return decisionID, http.StatusAccepted, nil
}
