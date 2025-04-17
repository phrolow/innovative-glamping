package services_test

import (
	"testing"
	"time"

	"phrolow/innovative_glamping/models"
	"phrolow/innovative_glamping/services"
)

func TestCheckAvailability(t *testing.T) {
	room := models.Room{ID: 1, IsAvailable: true}
	bookings := []models.Booking{
		{RoomID: 1, StartDate: parseDate("2025-04-10"), EndDate: parseDate("2025-04-15")},
	}

	tests := []struct {
		startDate string
		endDate   string
		expected  bool
	}{
		{"2025-04-05", "2025-04-09", true},  // Before existing booking
		{"2025-04-12", "2025-04-14", false}, // Overlapping existing booking
		{"2025-04-16", "2025-04-20", true},  // After existing booking
	}

	for _, test := range tests {
		start := parseDate(test.startDate)
		end := parseDate(test.endDate)
		result := services.CheckAvailability(room, bookings, start, end)
		if result != test.expected {
			t.Errorf("CheckAvailability(%s, %s) = %v; want %v", test.startDate, test.endDate, result, test.expected)
		}
	}
}

func TestBookRoom(t *testing.T) {
	room := models.Room{ID: 1, IsAvailable: true}
	bookings := []models.Booking{}

	err := services.BookRoom(&room, &bookings, "John Doe", parseDate("2025-04-10"), parseDate("2025-04-15"))
	if err != nil {
		t.Errorf("BookRoom() failed: %v", err)
	}

	// Attempt to book overlapping dates
	err = services.BookRoom(&room, &bookings, "Jane Doe", parseDate("2025-04-12"), parseDate("2025-04-18"))
	if err == nil {
		t.Errorf("BookRoom() should have failed for overlapping dates")
	}
}

func parseDate(date string) time.Time {
	parsed, _ := time.Parse("2006-01-02", date)
	return parsed
}
