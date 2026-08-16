package service

import (
	"studyroom/internal/model"
	"studyroom/internal/store"
)

type Service struct{ store *store.Store }

func New(s *store.Store) *Service { return &Service{store: s} }

func (s *Service) Book(id string) error {
	seat, _ := s.store.GetSeat(id)
	seat.Status = model.StatusBooked
	return nil
}
