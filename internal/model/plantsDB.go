package model

type GlossaryEntry struct {
    ID         uint   `gorm:"primaryKey"`
    Term       string
    Definition string
}