package database

import (
	"log"
	"reflect"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func BeforeCreate(model any, tx *gorm.DB) (err error) {
	val := reflect.ValueOf(model).Elem()
	idField := val.FieldByName("ID")

	if idField.IsValid() && idField.CanSet() && idField.Kind() == reflect.String {
		idField.SetString(uuid.New().String())
	}
	return nil
}

func InitDB(url string) {
	var err error
	db, err = gorm.Open(postgres.Open(url))
	if err != nil {
		log.Fatal("Failed to connect database!")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(0)

	log.Println("DB initialized with connection pool")
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get sql.DB for closing:", err)
		return
	}
	sqlDB.Close()
	log.Println("DB connection closed")
}
