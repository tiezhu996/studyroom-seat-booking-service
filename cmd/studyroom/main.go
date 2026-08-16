package main

import (
	"fmt"

	"studyroom/internal/model"
	"studyroom/internal/service"
	"studyroom/internal/store"
	"studyroom/internal/worker"
)

func main() {
	st := store.New()
	_ = st.AddSeat(&model.Seat{ID: "s1", Status: model.StatusFree})
	svc := service.New(st)
	_ = svc.Book("s1")
	fmt.Println("booked=", worker.New(st).CountBooked())
}
