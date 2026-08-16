package worker

import (
	"testing"

	"studyroom/internal/model"
	"studyroom/internal/store"
)

func TestCountBooked(t *testing.T) {
	s := store.New()
	_ = s.AddSeat(&model.Seat{ID: "b1", Status: model.StatusBooked})
	_ = s.AddSeat(&model.Seat{ID: "b2", Status: model.StatusBooked})
	_ = s.AddSeat(&model.Seat{ID: "f1", Status: model.StatusFree})
	if got := New(s).CountBooked(); got != 2 {
		t.Fatalf("booked=%d want 2", got)
	}
}
