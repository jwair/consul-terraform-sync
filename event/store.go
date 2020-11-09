package event

import (
	"errors"
	"fmt"
	"sync"
)

const defaultEventCountLimit = 5

// Store stores events
type Store struct {
	*sync.RWMutex

	events map[string][]*Event // taskname => events
	limit  int
}

// NewStore returns a new store
func NewStore() *Store {
	return &Store{
		RWMutex: &sync.RWMutex{},
		events:  make(map[string][]*Event),
		limit:   defaultEventCountLimit,
	}
}

// Add adds an event and manages the limit of number of events stored per task.
func (s *Store) Add(e *Event) error {
	if e == nil {
		return errors.New("error adding event: nil event")
	}
	if e.TaskName == "" {
		return fmt.Errorf("error adding event: taskname cannot be empty %s", e.GoString())
	}

	s.Lock()
	defer s.Unlock()

	events := s.events[e.TaskName]
	events = append([]*Event{e}, events...) // prepend
	if len(events) > s.limit {
		events = events[:len(events)-1]
	}
	s.events[e.TaskName] = events
	return nil
}

// Read returns events given a task name. Returned events are ordered by
// decending end time
func (s *Store) Read(taskName string) []Event {
	s.RLock()
	defer s.RUnlock()

	events := make([]Event, len(s.events[taskName]))
	for ix, event := range s.events[taskName] {
		events[ix] = *event
	}
	return events
}
