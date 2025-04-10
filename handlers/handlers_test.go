package handlers_test

import (
	"bytes"
	"encoding/json"
	"innovative_glamping/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetRooms(t *testing.T) {
	req := httptest.NewRequest("GET", "/rooms", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.GetRooms)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("GetRooms() status = %d; want %d", rr.Code, http.StatusOK)
	}

	var rooms []map[string]interface{}
	json.NewDecoder(rr.Body).Decode(&rooms)
	if len(rooms) == 0 {
		t.Errorf("GetRooms() returned no rooms; want > 0")
	}
}

func TestBookRoomHandler(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/rooms/{id}/book", handlers.BookRoom).Methods("POST")

	payload := map[string]string{
		"startDate": "2025-04-10",
		"endDate":   "2025-04-15",
		"customer":  "John Doe",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/rooms/1/book", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("BookRoomHandler() status = %d; want %d", rr.Code, http.StatusOK)
	}

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	if response["message"] != "Room booked successfully" {
		t.Errorf("BookRoomHandler() message = %s; want %s", response["message"], "Room booked successfully")
	}
}
