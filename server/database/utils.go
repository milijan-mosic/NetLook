package database

import (
	"log"
	"reflect"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	models "github.com/milijan-mosic/net-look/server/models"
)

func BeforeCreate(model interface{}, tx *gorm.DB) (err error) {
	val := reflect.ValueOf(model).Elem()
	idField := val.FieldByName("ID")

	if idField.IsValid() && idField.CanSet() && idField.Kind() == reflect.String {
		idField.SetString(uuid.New().String())
	}
	return nil
}

func OpenDBConnection(url string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(url))
	if err != nil {
		log.Fatal("Failed to connect database!")
	}

	log.Println("Database connection established!")
	return db
}

func closeDB(db *gorm.DB) error {
	dbConnection, err := db.DB()
	if err != nil {
		return err
	}
	return dbConnection.Close()
}

func CloseDBConnection(db *gorm.DB, debug bool) {
	if err := closeDB(db); err != nil {
		log.Fatalf("Error closing the database connection: %v", err)
	}

	if debug {
		log.Println("Database connection closed successfully!")
	}
}

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

func FindRootAgent(db *gorm.DB) []models.Agent {
	var allAgents []models.Agent
	if err := db.Where("name = ?", "Test").Find(&allAgents).Error; err != nil {
		log.Println("Failed to query agents by name: ", err)
	}

	if len(allAgents) > 0 {
		log.Println("Agent already exists: ", allAgents)
	}
	return allAgents
}

func SeedDatabase(db *gorm.DB) {
	allAgents := FindRootAgent(db)

	if len(allAgents) == 0 {
		agent := models.Agent{Name: "Test", UpdateInterval: 1.00}
		CreateModel(db, &agent, false)
	}
}
