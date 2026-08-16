package model

import "testing"

func TestIsBookedNilSafe(t *testing.T) {
	var s *Seat
	if s.IsBooked() {
		t.Fatal("nil seat should not be booked")
	}
	if got := (&Seat{Status: StatusBooked}).IsBooked(); !got {
		t.Fatal("booked seat should be booked")
	}
}
