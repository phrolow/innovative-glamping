package database

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1111"
	dbname   = "glamping"
)

// Connect initializes the database connection
func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// WithTransaction executes a set of database operations within a transaction
func WithTransaction(db *gorm.DB, txFunc func(tx *gorm.DB) error) error {
	tx := db.Begin() // Start the transaction

	if tx.Error != nil {
		return tx.Error
	}

	// Run the provided function within the transaction
	if err := txFunc(tx); err != nil {
		tx.Rollback() // Rollback the transaction on error
		return err
	}

	// Commit the transaction if everything is successful
	return tx.Commit().Error
}
