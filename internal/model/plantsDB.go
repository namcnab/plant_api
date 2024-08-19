package model

type Glossary struct {
    ID         int   `gorm:"primaryKey"`
    Term       string `gorm:"column:term"`
    Definition string `gorm:"column:definition"`
}

// TableName method to explicitly set the table name
func (Glossary) TableName() string {
    return "glossary"
}

// set schema name
func (Glossary) SchemaName() string {
    return "public"
}
