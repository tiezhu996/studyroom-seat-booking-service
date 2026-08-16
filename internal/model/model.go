package model

const (
	StatusFree   = "free"
	StatusBooked = "booked"
)

type Seat struct {
	ID     string
	Status string
}

func (s *Seat) IsBooked() bool {
	return s.Status == StatusBooked
}
