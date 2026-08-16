package service

import (
	"fmt"

	"studyroom/internal/model"
	"studyroom/internal/store"
)

type Service struct{ store *store.Store }

func New(s *store.Store) *Service { return &Service{store: s} }

func (s *Service) Book(id string) error {
	seat, ok := s.store.GetSeat(id)
	if !ok {
		return fmt.Errorf("seat not found")
	}
	seat.Status = model.StatusBooked
	return nil
}
