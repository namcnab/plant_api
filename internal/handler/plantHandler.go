package handler

import (
	dbHandler "github.com/namcnab/plant_api/internal/database"
	m "github.com/namcnab/plant_api/internal/model"
	"gorm.io/gorm"
)

func GetGlossary(db *gorm.DB) ([]m.Glossary, error) {
	dbResp, err := dbHandler.GetAllGlossaryEntries(db)

	if err != nil {
		return nil, err
	}

	return dbResp, nil
}