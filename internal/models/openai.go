package models

type OpenAIRequest struct {
	Prompt string `json:"prompt"`
}

type OpenAIResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}
