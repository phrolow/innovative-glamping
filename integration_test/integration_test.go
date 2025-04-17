package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"innovative_glamping/handlers"
)

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Public Routes
	router.HandleFunc("/rooms", handlers.GetRooms).Methods("GET")
	router.HandleFunc("/rooms/{id}/availability", handlers.CheckRoomAvailability).Methods("GET")

	// Protected Routes (simulate authentication if needed)
	router.HandleFunc("/rooms/{id}/book", handlers.BookRoom).Methods("POST")
	router.HandleFunc("/cancel", handlers.CancelBooking).Methods("POST")

	return router
}

func TestBookingIntegration(t *testing.T) {
	router := setupRouter()

	// Test successful room booking
	t.Run("Success Booking", func(t *testing.T) {
		payload := map[string]string{
			"startDate": "2025-04-20",
			"endDate":   "2025-04-25",
			"customer":  "John Doe",
		}
		body, _ := json.Marshal(payload)

		req := httptest.NewRequest("POST", "/rooms/1/book", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rr.Code)
		}

		var response map[string]string
		_ = json.NewDecoder(rr.Body).Decode(&response)
		if response["message"] != "Room booked successfully" {
			t.Errorf("Expected success message, got %s", response["message"])
		}
	})

	// Test checking room availability
	t.Run("Check Room Availability", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/rooms/1/availability?start=2025-04-20&end=2025-04-25", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rr.Code)
		}

		var response map[string]bool
		_ = json.NewDecoder(rr.Body).Decode(&response)
		if response["available"] {
			t.Errorf("Expected room to be unavailable, got %v", response["available"])
		}
	})

	// Test booking an already booked room
	t.Run("Double Booking", func(t *testing.T) {
		payload := map[string]string{
			"startDate": "2025-04-22",
			"endDate":   "2025-04-24",
			"customer":  "Jane Doe",
		}
		body, _ := json.Marshal(payload)

		req := httptest.NewRequest("POST", "/rooms/1/book", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", rr.Code)
		}

		var response map[string]string
		_ = json.NewDecoder(rr.Body).Decode(&response)
		if response["message"] != "room is not available for the selected dates" {
			t.Errorf("Expected error message, got %s", response["message"])
		}
	})
}
