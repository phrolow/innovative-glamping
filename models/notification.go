package models

import "time"

// Notification represents a notification message
type Notification struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}
