package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"innovative_glamping/handlers"
	"innovative_glamping/middleware"
)

func main() {
	r := mux.NewRouter()

	// Apply error handler middleware
	r.Use(handlers.ErrorHandler)

	// Public Routes
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected Routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.Authenticate)
	protected.HandleFunc("/rooms", handlers.GetRooms).Methods("GET")
	protected.HandleFunc("/rooms/{id}/availability", handlers.CheckRoomAvailability).Methods("GET")
	protected.HandleFunc("/rooms/{id}/book", handlers.BookRoom).Methods("POST")
	protected.HandleFunc("/cancel", handlers.CancelBooking).Methods("POST")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
