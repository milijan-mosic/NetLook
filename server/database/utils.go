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

func OpenDB(url string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(url))
	if err != nil {
		log.Fatal("Failed to connect database!")
	}

	log.Println("Database connection established!")
	return db
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

func SeedDatabase(db *gorm.DB, debug bool) {
	var allAgents []models.Agent
	if err := db.Where("name = ?", "Test").Find(&allAgents).Error; err != nil {
		log.Fatal("Failed to query agents by name: ", err)
	}

	if len(allAgents) > 0 {
		log.Println("Agent already exists:", allAgents)
		return
	}

	agent := models.Agent{Name: "Test", UpdateInterval: 1.00}
	if result := db.Create(&agent); result.Error != nil {
		log.Println("Error inserting agent:", result.Error)
	} else {
		log.Println("Inserted test agent:", agent)
	}
}

// // 5. Update the user's age
// db.Model(&queriedUser).Update("Age", 30)
// fmt.Println("Updated user's age")

// // 6. Delete the user
// db.Delete(&queriedUser)
// fmt.Println("Deleted user")
