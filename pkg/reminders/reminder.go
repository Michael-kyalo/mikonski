package reminders

import (
	"errors"
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

func (s *Scheduler) Set(description string, t time.Time) error {
	if t.Before(time.Now()) {
		return errors.New("time must be in the future")
	}
	s.reminders = append(s.reminders, Reminder{Description: description, Time: t})
	return nil
}

func (s *Scheduler) List() []Reminder {
	return s.reminders
}

func (s *Scheduler) Clear() {
	s.reminders = []Reminder{}
}
