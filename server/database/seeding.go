package database

import (
	"gorm.io/gorm"

	models "github.com/milijan-mosic/net-look/server/models"
)

func SeedDatabase(db *gorm.DB) {
	allAgents := FindRootAgent(db)

	if len(allAgents) == 0 {
		agent := models.Agent{Name: "Test", UpdateInterval: 1.00}
		CreateModel(db, &agent, false)
	}
}
