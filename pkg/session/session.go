package session

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Michael-kyalo/mikonski/pkg/logging"
	"go.uber.org/zap"
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

const sessionFile = "session.json" // File to store session data

// NewSession initializes a new session, loading data from a file if available.
func NewSession() *Session {
	session := &Session{history: []QuestionResponse{}}
	if err := session.LoadFromFile(); err != nil && !errors.Is(err, os.ErrNotExist) {
		logging.GetLogger().Error("failed to load session from file", zap.Error(err))
	}
	return session
}

// SaveToFile saves the session history to a file.
func (s *Session) SaveToFile() error {
	logger := logging.GetLogger()

	file, err := os.Create(sessionFile)
	if err != nil {
		logger.Error("failed to create session file", zap.Error(err))
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(s.history); err != nil {
		logger.Error("failed to encode session history to file", zap.Error(err))
		return err
	}

	logger.Info("session saved to file", zap.String("file", sessionFile))
	return nil
}

// LoadFromFile loads the session history from a file.
func (s *Session) LoadFromFile() error {
	logger := logging.GetLogger()

	file, err := os.Open(sessionFile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.history); err != nil {
		logger.Error("failed to decode session history from file", zap.Error(err))
		return err
	}

	logger.Info("session loaded from file", zap.String("file", sessionFile))
	return nil
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
