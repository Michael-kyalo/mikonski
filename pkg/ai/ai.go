package ai

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
)

type Client interface {
	Ask(question string) (string, error)
}

type OpenAIClient struct {
	apiKey string
	model  string
	client *openai.Client
}

// NewOpenAIClient initializes a new OpenAIClient.
func NewOpenAIClient(apiKey, model string) *OpenAIClient {
	return &OpenAIClient{
		apiKey: apiKey,
		model:  model,
		client: openai.NewClient(apiKey),
	}
}

// Ask sends a question to the OpenAI API and returns its response.
func (o *OpenAIClient) Ask(question string) (string, error) {
	if question == "" {
		return "", errors.New("question cannot be empty")
	}

	ctx := context.Background()
	resp, err := o.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: o.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	})

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}

// mock for testing
type MockClient struct {
	Response string
}

func (c MockClient) Ask(question string) (string, error) {
	return c.Response, nil
}
