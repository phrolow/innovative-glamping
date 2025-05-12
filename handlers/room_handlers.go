package handlers

import (
	"encoding/json"
	"fmt"
	"innovative_glamping/models"
	"innovative_glamping/services"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var rooms = []models.Room{
	{ID: 1, Type: "Deluxe", IsAvailable: true},
	{ID: 2, Type: "Standard", IsAvailable: true},
}

var bookings = []models.Booking{}

// GetRooms retrieves all rooms
func GetRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}

// CheckRoomAvailability checks availability for a room in a given date range
func CheckRoomAvailability(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the room
	var room models.Room
	for _, r := range rooms {
		if r.ID == id {
			room = r
			break
		}
	}
	if room.ID == 0 {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	// Parse query parameters
	startDate, _ := time.Parse("2006-01-02", r.URL.Query().Get("start"))
	endDate, _ := time.Parse("2006-01-02", r.URL.Query().Get("end"))

	// Check availability
	isAvailable := services.CheckAvailability(room, bookings, startDate, endDate)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"available": isAvailable})
}

// BookRoom books a room if available
func BookRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the room
	var room *models.Room
	for i := range rooms {
		if rooms[i].ID == id {
			room = &rooms[i]
			break
		}
	}
	if room == nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	// Parse request body
	var bookingRequest struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Customer  string `json:"customer"`
	}
	_ = json.NewDecoder(r.Body).Decode(&bookingRequest)

	// Convert dates
	startDate, _ := time.Parse("2006-01-02", bookingRequest.StartDate)
	endDate, _ := time.Parse("2006-01-02", bookingRequest.EndDate)

	// Book the room
	err := services.BookRoom(room, &bookings, bookingRequest.Customer, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Room booked successfully"})

	services.SendNotification(models.Notification{
		Type:    "Booking",
		Message: fmt.Sprintf("Room %d has been successfully booked by %s", room.ID, bookingRequest.Customer),
	})
}

func CancelBooking(w http.ResponseWriter, r *http.Request) {
	var cancelRequest struct {
		ID string `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&cancelRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	cancelId, _ := strconv.Atoi(cancelRequest.ID)

	// Find and remove the booking
	for i, booking := range bookings {
		log.Println(booking.ID)
		if booking.ID == cancelId {
			bookings = append(bookings[:i], bookings[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Booking canceled successfully"})
			return
		}
	}

	http.Error(w, "Booking not found", http.StatusNotFound)
}
