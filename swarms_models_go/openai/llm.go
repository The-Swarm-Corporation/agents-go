package openai

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
)

type OpenAIClient struct {
	APIKey string
	APIURL string
	Model  string
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Choice struct {
	Message Message `json:"message"`
}

type ChatCompletionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

func NewClient(apiKey, model string) *OpenAIClient {
	return &OpenAIClient{
		APIKey: apiKey,
		APIURL: "https://api.openai.com/v1/chat/completions",
		Model:  model,
	}
}

func (client *OpenAIClient) CreateChatCompletion(prompt string) (string, error) {
	req := ChatCompletionRequest{
		Model: client.Model,
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	restyClient := resty.New()
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+client.APIKey).
		SetBody(req).
		Post(client.APIURL)

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", errors.New(resp.String())
	}

	var chatCompletionResponse ChatCompletionResponse
	if err := json.Unmarshal(resp.Body(), &chatCompletionResponse); err != nil {
		return "", err
	}

	if len(chatCompletionResponse.Choices) == 0 {
		return "", errors.New("no choices returned")
	}

	return chatCompletionResponse.Choices[0].Message.Content, nil
}
