package services_test

import (
	"innovative_glamping/models"
	"innovative_glamping/services"
	"testing"
)

func TestNotificationService(t *testing.T) {
	go services.StartNotificationListener()

	notification := models.Notification{
		ID:      1,
		Type:    "Test",
		Message: "This is a test notification",
	}

	services.SendNotification(notification)
}
