package service

import (
	"testing"

	"studyroom/internal/store"
)

func TestBookMissingReturnsError(t *testing.T) {
	s := New(store.New())
	if err := s.Book("missing"); err == nil {
		t.Fatal("expected error for missing seat")
	}
}
