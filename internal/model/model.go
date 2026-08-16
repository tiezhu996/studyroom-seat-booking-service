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
	if s == nil {
		return false
	}
	return s.Status == StatusBooked
}
