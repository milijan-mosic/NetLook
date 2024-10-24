package database

import (
	"log"
	"reflect"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ID struct {
	ID string `gorm:"type:uuid;primaryKey" json:"id"`
}

func (base *ID) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New().String()
	return
}

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

// // 3. Create a new record
// user := User{Name: "Alice", Age: 25}
// result := db.Create(&user) // Pass pointer of data to create
// if result.Error != nil {
// 	fmt.Println(result.Error)
// }
// fmt.Println("Inserted User ID:", user.ID)

// // 4. Query a user
// var queriedUser User
// db.First(&queriedUser, 1) // Find the user with primary key 1
// fmt.Printf("Queried User: %+v\n", queriedUser)

// // 5. Update the user's age
// db.Model(&queriedUser).Update("Age", 30)
// fmt.Println("Updated user's age")

// // 6. Delete the user
// db.Delete(&queriedUser)
// fmt.Println("Deleted user")
