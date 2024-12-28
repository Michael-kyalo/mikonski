package reminders

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/Michael-kyalo/mikonski/pkg/logging"
	"go.uber.org/zap"
)

// Reminder represents a single reminder entry.
type Reminder struct {
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}

// Scheduler manages reminders.
type Scheduler struct {
	reminders []Reminder
}

const remindersFile = "reminders.json" // File to store reminders

// NewScheduler initializes a new Scheduler.// NewScheduler initializes a new Scheduler, loading data from a file if available.
func NewScheduler() *Scheduler {
	scheduler := &Scheduler{reminders: []Reminder{}}
	if err := scheduler.LoadFromFile(); err != nil && !errors.Is(err, os.ErrNotExist) {
		logging.GetLogger().Error("failed to load reminders from file", zap.Error(err))
	}
	return scheduler
}

// SaveToFile saves the reminders to a file.
func (s *Scheduler) SaveToFile() error {
	logger := logging.GetLogger()

	file, err := os.Create(remindersFile)
	if err != nil {
		logger.Error("failed to create reminders file", zap.Error(err))
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(s.reminders); err != nil {
		logger.Error("failed to encode reminders to file", zap.Error(err))
		return err
	}

	return nil
}

// LoadFromFile loads reminders from a file.
func (s *Scheduler) LoadFromFile() error {
	logger := logging.GetLogger()

	file, err := os.Open(remindersFile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.reminders); err != nil {
		logger.Error("failed to decode reminders from file", zap.Error(err))
		return err
	}

	return nil
}

// Set adds a new reminder to the scheduler.
func (s *Scheduler) Set(description string, t time.Time) error {
	if t.Before(time.Now()) {
		return errors.New("time must be in the future")
	}
	s.reminders = append(s.reminders, Reminder{Description: description, Time: t})
	return nil
}

// ExportReminders saves all reminders to a JSON file.
func (s *Scheduler) ExportReminders(filename string) error {
	logger := logging.GetLogger()

	// Log the number of reminders
	logger.Info("exporting reminders", zap.Int("reminder_count", len(s.reminders)))

	// Log each reminder for debugging
	for i, reminder := range s.reminders {
		logger.Debug("reminder entry",
			zap.Int("index", i),
			zap.String("description", reminder.Description),
			zap.Time("time", reminder.Time),
		)
	}

	// Create the file for exporting reminders
	file, err := os.Create(filename)
	if err != nil {
		logger.Error("failed to create export file", zap.Error(err))
		return err
	}
	defer file.Close()

	// Write the reminders to the file in JSON format
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(s.reminders); err != nil {
		logger.Error("failed to encode reminders to JSON", zap.Error(err))
		return err
	}

	return nil
}

// List returns all the reminders in the scheduler.
func (s *Scheduler) List() []Reminder {
	return s.reminders
}

// Clear removes all reminders from the scheduler.
func (s *Scheduler) Clear() {
	s.reminders = []Reminder{}
}
