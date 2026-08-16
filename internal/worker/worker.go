package worker

import "studyroom/internal/store"

type Worker struct{ store *store.Store }

func New(s *store.Store) *Worker { return &Worker{store: s} }

func (w *Worker) CountBooked() int {
	count := 0
	for _, seat := range w.store.AllSeats() {
		if !seat.IsBooked() {
			count++
		}
	}
	return count
}
