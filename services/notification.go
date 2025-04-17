package services

import (
	"fmt"
	"innovative_glamping/models"
	"time"
)

var notificationChannel = make(chan models.Notification, 10)

// StartNotificationListener starts listening for notifications
func StartNotificationListener() {
	go func() {
		for notification := range notificationChannel {
			// Simulate processing the notification
			fmt.Printf("Processing notification: %v\n", notification)
		}
	}()
}

// SendNotification sends a notification to the channel
func SendNotification(notification models.Notification) {
	notification.CreatedAt = time.Now()
	notificationChannel <- notification
}
