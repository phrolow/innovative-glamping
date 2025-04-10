package services

import (
	"errors"
	"time"

	"phrolow/innovative-glamping/models"
)

// CheckAvailability checks if a room is available for a given date range
func CheckAvailability(room models.Room, bookings []models.Booking, startDate, endDate time.Time) bool {
	if !room.IsAvailable {
		return false
	}

	for _, booking := range bookings {
		if booking.RoomID == room.ID {
			// Check if dates overlap
			if startDate.Before(booking.EndDate) && endDate.After(booking.StartDate) {
				return false
			}
		}
	}
	return true
}

// BookRoom books a room if available
func BookRoom(room *models.Room, bookings *[]models.Booking, customer string, startDate, endDate time.Time) error {
	if !CheckAvailability(*room, *bookings, startDate, endDate) {
		return errors.New("room is not available for the selected dates")
	}

	// Create a new booking
	newBooking := models.Booking{
		ID:        len(*bookings) + 1, // Simple ID generation
		RoomID:    room.ID,
		StartDate: startDate,
		EndDate:   endDate,
		Customer:  customer,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	*bookings = append(*bookings, newBooking)
	return nil
}