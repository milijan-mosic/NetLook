package database

import (
	"log"
	"reflect"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(url))
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
