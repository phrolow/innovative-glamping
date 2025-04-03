package main

import (
	"innovative_glamping/database"
	"innovative_glamping/models"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Выполнение миграций
	err = db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatalf("Ошибка выполнения миграций: %v", err)
	}

	log.Println("Миграции успешно применены!")
}
