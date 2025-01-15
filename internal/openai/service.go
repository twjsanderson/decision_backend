package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type OpenAIService struct {
	client *openai.Client
}

func NewOpenAIService(apiKey string) *OpenAIService {
	client := openai.NewClient(apiKey)
	return &OpenAIService{client: client}
}

func (s *OpenAIService) GetChatGPTResponse(prompt string, ctx string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: ctx,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	resp, err := s.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("failed to get response from OpenAI: %v", err)
	}

	// Extract the response content
	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from OpenAI")
}

func GetChatGPTResponse(apiKey string, prompt string, ctx string) (string, error) {
	openAIService := NewOpenAIService(apiKey)

	response, err := openAIService.GetChatGPTResponse(prompt, ctx)
	if err != nil {
		return "", fmt.Errorf("error while getting response: %v", err)
	}

	return response, nil
}
