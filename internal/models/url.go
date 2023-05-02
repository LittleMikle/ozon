package models

import "gorm.io/gorm"

type URLs struct {
	URL     *string `json:"url"`
	TinyURL *string `gorm:"primaryKey;autoIncrement:false" json:"tiny_url"`
}

func MigrateURLs(db *gorm.DB) error {
	err := db.AutoMigrate(&URLs{})
	return err
}
