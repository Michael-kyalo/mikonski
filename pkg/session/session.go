package session

import (
	"encoding/json"
	"os"
)

// QuestionResponse represents a question and its response.
type QuestionResponse struct {
	Question string `json:"question"`
	Response string `json:"response"`
}

// Session manages context during a session.
type Session struct {
	history []QuestionResponse
}

// NewSession initializes a new session.
func NewSession() *Session {
	return &Session{history: []QuestionResponse{}}
}

// AddContext updates the session context and history.
func (s *Session) AddContext(question, response string) {
	s.history = append(s.history, QuestionResponse{
		Question: question,
		Response: response,
	})
}

// ExportHistory saves the session history to a JSON file.
func (s *Session) ExportHistory(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(s.history)
}

// GetContext retrieves the most recent response in the session history.
func (s *Session) GetContext() string {
	if len(s.history) == 0 {
		return ""
	}
	return s.history[len(s.history)-1].Response
}

// ClearContext clears the session context and history.
func (s *Session) ClearContext() {
	s.history = []QuestionResponse{}
}
