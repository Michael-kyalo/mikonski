package tests

import (
	"testing"

	"github.com/Michael-kyalo/mikonski/pkg/ai"
)

/**
Ensure the ask command accepts input and invokes the AI client.
Mock the AI client to return dummy responses for testing.
*/

func TestAskCommand(t *testing.T) {
	// Mock AI Client
	mockClient := ai.MockClient{
		Response: "This is a test response",
	}

	// Simulate user input
	question := "What is the capital of France?"
	response, err := mockClient.Ask(question)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "This is a test response"
	if response != expected {
		t.Errorf("Expected %q, got %q", expected, response)
	}
}
