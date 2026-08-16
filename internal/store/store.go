package store

import (
	"fmt"
	"sync"

	"studyroom/internal/model"
)

type Store struct {
	mu    sync.Mutex
	seats map[string]*model.Seat
}

func New() *Store {
	return &Store{seats: make(map[string]*model.Seat)}
}

func (s *Store) AddSeat(seat *model.Seat) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if seat == nil {
		return fmt.Errorf("nil seat")
	}
	if seat.ID == "" {
		return fmt.Errorf("seat id required")
	}
	s.seats[seat.ID] = seat
	return nil
}

func (s *Store) GetSeat(id string) (*model.Seat, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	seat, ok := s.seats[id]
	return seat, ok
}

func (s *Store) AllSeats() []*model.Seat {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]*model.Seat, 0, len(s.seats))
	for _, seat := range s.seats {
		out = append(out, seat)
	}
	return out
}
