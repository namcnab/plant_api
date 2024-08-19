package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	m "github.com/namcnab/plant_api/internal/model" // Import the package that contains the definition of GlossaryEntry
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, errors.New("failed to load environment variables")
	}

    // Retrieve database configuration from environment variables
	host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")
	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

    // Connect to the database using GORM
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, errors.New("failed to connect to database")
    }

    return db, nil
}

func GetAllGlossaryEntries(db *gorm.DB) ([]m.Glossary, error) {
    var glossary []m.Glossary
    // Retrieve all entries from the public.glossary table
    result := db.Find(&glossary)

    if result.Error != nil {
        return nil, errors.New("failed to retrieve glossary entries")
    }

    return glossary, nil
}
