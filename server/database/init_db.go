package database

import (
	"fmt"

	models "github.com/milijan-mosic/net-look/server/models"
)

func GetDBUrl() string {
	username := GetEnvVariable("POSTGRES_USER")
	password := GetEnvVariable("POSTGRES_PASSWORD")
	dbName := GetEnvVariable("POSTGRES_DB")

	return fmt.Sprintf("postgres://%s:%s@database:5432/%s", username, password, dbName)
}

func InitializeDatabase() {
	url := GetDBUrl()
	InitDB(url)
	MigrateModels(db, models.AllModels)
	SeedDatabase(db)
}
