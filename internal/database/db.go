package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/namcnab/plant_api/internal/model"
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

func CreateGlossaryEntry(db *gorm.DB, entry m.Glossary) error {
    // Insert a new entry into the public.glossary table

    if err := db.Where("term = ?", entry.Term).First(&entry).Error; err == nil {
        UpdateGlossaryEntry(db, entry)
    } else {
        result := db.Create(&entry)

        if result.Error != nil {
            return errors.New("failed to add glossary entry")
        }
    }

    return nil
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

func UpdateGlossaryEntry(db *gorm.DB, entry m.Glossary) error {
    // Update an existing entry in the public.glossary table
    definition := entry.Definition

    if err := db.Where("term = ?", entry.Term).First(&entry).Error; err != nil {
        return errors.New("entry not found")
    }
   
    result := db.Model(&entry).Where("term = ?", entry.Term).Update("definition", definition)
    
    if result.Error != nil {
        return errors.New("failed to update glossary entry")
    }

    return nil
}

func DeleteGlossaryTerm(db *gorm.DB, term string) error {
    // Delete an existing entry from the public.glossary table
    if err := db.Where("term = ?", term).First(&model.Glossary{}).Error; err != nil {
        return errors.New("entry not found")
    }

    result := db.Model(&model.Glossary{}).Where("term = ?", term).Delete(&model.Glossary{})

    if result.Error != nil {
        return errors.New("failed to delete glossary entry")
    }

    return nil
}