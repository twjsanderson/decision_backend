package models

import "encoding/json"

type DecisionFields struct {
	ParetoAnalysis      *ParetoAnalysis           `json:"pareto_analysis,omitempty"`
	SwotAnalysis        *SwotAnalysis             `json:"swot_analysis,omitempty"`
	BayesianDecision    *BayesianDecision         `json:"bayesian_decision,omitempty"`
	DecisionTree        *DecisionTree             `json:"decision_tree,omitempty"`
	AHP                 *AnalyticHierarchyProcess `json:"ahp,omitempty"`
	FirstPrinciples     *FirstPrinciples          `json:"first_principles,omitempty"`
	FuzzyLogic          *FuzzyLogicDecisionMaking `json:"fuzzy_logic,omitempty"`
	CostBenefitAnalysis *CostBenefitAnalysis      `json:"cost_benefit_analysis,omitempty"`
}

type Decision struct {
	User
	DecisionId int    `json:"decisionId"`
	Title      string `json:"title"`
	Problem    string `json:"problem"`
	DecisionFields
}

type DecisionResponse struct {
	DecisionId int    `json:"decisionId"`
	Title      string `json:"title"`
	Problem    string `json:"problem"`
	UserId     string `json:"userId"`
	DecisionFields
}

type ParetoAnalysis struct {
	PossibleCauses  []string `json:"possible_causes"`
	ImpactMeasures  []string `json:"impact_measures"`
	ExpectedOutcome string   `json:"expected_outcome"`
}

type SwotAnalysis struct {
	Strengths     []string `json:"strengths"`
	Weaknesses    []string `json:"weaknesses"`
	Opportunities []string `json:"opportunities"`
	Threats       []string `json:"threats"`
}

type BayesianDecision struct {
	Hypothesis                string  `json:"hypothesis"`
	PriorProbability          float64 `json:"prior_probability"`
	Evidence                  string  `json:"evidence"`
	LikelihoodHypothesisTrue  float64 `json:"likelihood_hypothesis_true"`
	LikelihoodHypothesisFalse float64 `json:"likelihood_hypothesis_false"`
	ExpectedOutcome           string  `json:"expected_outcome"`
}

type DecisionTree struct {
	Options json.RawMessage `json:"options"`
}

type AnalyticHierarchyProcess struct {
	Criteria     json.RawMessage `json:"criteria"`
	Alternatives json.RawMessage `json:"alternatives"`
}

type FirstPrinciples struct {
	Assumptions           json.RawMessage `json:"assumptions"`
	FundamentalFacts      json.RawMessage `json:"fundamental_facts"`
	ReconstructedSolution string          `json:"reconstructed_solution"`
}

type FuzzyLogicDecisionMaking struct {
	FuzzyVariables    json.RawMessage `json:"fuzzy_variables"`
	DecisionThreshold float64         `json:"decision_threshold"`
}

type CostBenefitAnalysis struct {
	Costs        json.RawMessage `json:"costs"`
	Benefits     json.RawMessage `json:"benefits"`
	DiscountRate float64         `json:"discount_rate"`
	TimeHorizon  int             `json:"time_horizon"`
}
