package models

import "time"

// Room represents a room in the glamping site
type Room struct {
	ID          int       `json:"id"`
	Type        string    `json:"type"`
	IsAvailable bool      `json:"isAvailable"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Booking represents a booking for a room
type Booking struct {
	ID        int       `json:"id"`
	RoomID    int       `json:"roomId"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Customer  string    `json:"customer"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}