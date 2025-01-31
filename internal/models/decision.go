package models

type Decision struct {
	User
	DecisionId int    `json:"decisionId"`
	Title      string `json:"title"`
	Problem    string `json:"problem"`
	// ChoiceType string `json:"choiceType"`
	// Justification string `json:"justification"`
	// IdealOutcome  string `json:"idealOutcome"`
	// MaxCost       string `json:"maxCost"`
	// RiskTolerance string `json:"riskTolerance"`
	// Timeline      string `json:"timeline"`
}

type DecisionResponse struct {
	DecisionId int    `json:"decisionId"`
	Title      string `json:"title"`
	Problem    string `json:"problem"`
	UserId     string `json:"userId"`
}
