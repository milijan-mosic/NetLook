package database

import (
	"log"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB, models []interface{}) error {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatal("Failed to migrate models: ", err)
			return err
		}
	}

	log.Println("Models migrated successfully!")
	return nil
}

func CreateModel[T any](db *gorm.DB, model T, debug bool) error {
	result := db.Create(&model)
	if result.Error != nil {
		log.Println("Error inserting model: ", result.Error)
		return result.Error
	}

	if debug {
		log.Println("Inserted model: ", model)
	}
	return nil
}

func CreateModels[T any](db *gorm.DB, models []T, debug bool) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&models).Error; err != nil {
		tx.Rollback()
		log.Println("Error inserting models: ", err)
		return err
	}

	tx.Commit()
	if debug {
		log.Println("Inserted models: ", models)
	}
	return nil
}
