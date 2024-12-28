package tests

import (
	"github.com/Michael-kyalo/mikonski/pkg/session"
	"testing"
)

func TestAddContext(t *testing.T) {
	sess := session.NewSession()

	// Add context
	sess.AddContext("What is Go?", "Go is a programming language.")

	// Retrieve context
	context := sess.GetContext()
	expected := "Go is a programming language."
	if context != expected {
		t.Errorf("Expected %q, got %q", expected, context)
	}
}

func TestClearContext(t *testing.T) {
	sess := session.NewSession()

	// Add and clear context
	sess.AddContext("What is Go?", "Go is a programming language.")
	sess.ClearContext()

	// Check if context is empty
	context := sess.GetContext()
	if context != "" {
		t.Errorf("Expected empty context, got %q", context)
	}
}
