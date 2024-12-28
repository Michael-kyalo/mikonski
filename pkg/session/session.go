package session

// Session manages context during a session.
type Session struct {
	context string
}

// NewSession initializes a new session.
func NewSession() *Session {
	return &Session{context: ""}
}

// AddContext updates the session context.
func (s *Session) AddContext(question, answer string) {
	s.context = answer // Store only the latest response as context.
}

// GetContext retrieves the current context.
func (s *Session) GetContext() string {
	return s.context
}

// ClearContext clears the session context.
func (s *Session) ClearContext() {
	s.context = ""
}
