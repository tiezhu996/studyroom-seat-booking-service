package store

import "testing"

func TestNewInitializesMap(t *testing.T) {
	s := New()
	if s.seats == nil {
		t.Fatal("seats map should be initialized")
	}
}

func TestGetSeatMissing(t *testing.T) {
	s := New()
	if seat, ok := s.GetSeat("missing"); ok || seat != nil {
		t.Fatalf("expected ok=false, seat=nil; got ok=%v seat=%v", ok, seat)
	}
}
