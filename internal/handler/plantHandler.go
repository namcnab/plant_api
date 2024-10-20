package handler

import (
	dbHandler "github.com/namcnab/plant_api/internal/database"
	m "github.com/namcnab/plant_api/internal/model"
	"gorm.io/gorm"
)

func CreateGlossaryEntry(db *gorm.DB, entry m.Glossary) error {
	err := dbHandler.CreateGlossaryEntry(db, entry)

	if err != nil {
		return err
	}

	return nil
} 

func GetGlossary(db *gorm.DB) ([]m.Glossary, error) {
	dbResp, err := dbHandler.GetAllGlossaryEntries(db)

	if err != nil {
		return nil, err
	}

	return dbResp, nil
}

func UpdateGlossaryEntry(db *gorm.DB, entry m.Glossary) error {
	err := dbHandler.UpdateGlossaryEntry(db, entry)

	if err != nil {
		return err
	}

	return nil
}

func DeleteGlossaryTerm(db *gorm.DB, term string) error {
	err := dbHandler.DeleteGlossaryTerm(db, term)

	if err != nil {
		return err
	}

	return nil
}