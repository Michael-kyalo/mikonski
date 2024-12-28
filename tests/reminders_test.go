package tests

import (
	"testing"
	"time"

	"github.com/Michael-kyalo/mikonski/pkg/reminders"
)

/**
reminder set: Ensure reminders are added with correct timestamps and descriptions.
reminder list: Ensure all reminders are displayed.
reminder clear: Ensure reminders can be cleared individually or entirely.
*/

func TestReminderSet(t *testing.T) {
	scheduler := reminders.NewScheduler()
	err := scheduler.Set("Test Reminder", time.Now().Add(1*time.Hour))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(scheduler.List()) != 1 {
		t.Errorf("Expected 1 reminder, got %d", len(scheduler.List()))
	}
}

func TestReminderClear(t *testing.T) {
	scheduler := reminders.NewScheduler()
	scheduler.Set("Test Reminder", time.Now().Add(1*time.Hour))
	scheduler.Clear()
	if len(scheduler.List()) != 0 {
		t.Errorf("Expected 0 reminders, got %d", len(scheduler.List()))
	}
}
