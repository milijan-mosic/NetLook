package database

import (
	"log"

	"gorm.io/gorm"

	models "github.com/milijan-mosic/net-look/server/models"
)

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
