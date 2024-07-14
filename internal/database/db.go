package database

import (
	"errors"
	"fmt"
	"os"

	m "github.com/namcnab/plant_api/model" // Import the models package
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializeDB() (*gorm.DB, error) {

    // Retrieve database configuration from environment variables
	host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")
	
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

    // Connect to the database using GORM
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, errors.New("failed to connect to database")
    }

    return db, nil
}
// q: how to resolve the error in the following function?
// a: import the models package 

func getAllGlossaryEntries(db *gorm.DB) ([]m.GlossaryEntry, error) {
    var entries []m.GlossaryEntry
    result := db.Find(&entries)
    if result.Error != nil {
        return nil, result.Error
    }
    return entries, nil
}
