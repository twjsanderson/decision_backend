package models

type Decision struct {
	User
	Title         string `json:"title"`
	ChoiceType    string `json:"choice_type"`
	Problem       string `json:"problem"`
	Justification string `json:"justification"`
	IdealOutcome  string `json:"ideal_outcome"`
	MaxCost       string `json:"max_cost"`
	RiskTolerance string `json:"risk_tolerance"`
	Timeline      string `json:"timeline"`
}
