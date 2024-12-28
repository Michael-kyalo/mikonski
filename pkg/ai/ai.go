package ai

import "errors"

type Client interface {
	Ask(question string) (string, error)
}

type OpenAIClient struct{}

func (c OpenAIClient) Ask(question string) (string, error) {
	// Call OpenAI API here
	return "", errors.New("not implemented")
}

type MockClient struct {
	Response string
}

func (c MockClient) Ask(question string) (string, error) {
	return c.Response, nil
}
