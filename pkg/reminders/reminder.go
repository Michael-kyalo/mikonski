package reminders

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Reminder struct {
	Description string
	Time        time.Time
}

type Scheduler struct {
	reminders []Reminder
}

func NewScheduler() *Scheduler {
	return &Scheduler{reminders: []Reminder{}}
}

// set schedules a new Reminder
func (s *Scheduler) Set(description string, t time.Time) error {
	if t.Before(time.Now()) {
		return errors.New("time must be in the future")
	}
	s.reminders = append(s.reminders, Reminder{Description: description, Time: t})
	return nil
}

// fetches a list of all the reminders
func (s *Scheduler) List() []Reminder {
	return s.reminders
}

// clear clears all the reminders
func (s *Scheduler) Clear() {
	s.reminders = []Reminder{}
}

// ExportReminders saves all reminders to a JSON file.
func (s *Scheduler) ExportReminders(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(s.reminders)
}
