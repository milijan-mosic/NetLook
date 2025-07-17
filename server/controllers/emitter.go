package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	database "github.com/milijan-mosic/net-look/server/database"
	models "github.com/milijan-mosic/net-look/server/models"
)

func ServeClient(w http.ResponseWriter, r *http.Request) {
	url := database.GetDBUrl()
	db := database.OpenDBConnection(url)

	allAgents := database.FindRootAgent(db)
	if len(allAgents) == 0 {
		http.Error(w, "No agents found", http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()
	time30SecondsAgo := currentTime.Add(-30 * time.Second).Unix()

	var data models.ClientPackage

	for _, agent := range allAgents {
		var agentData models.AgentData
		agentData.AgentName = agent.Name

		var cpuData []models.CPU
		if err := db.Where("agent_id = ? AND timestamp >= ?", agent.ID, time30SecondsAgo).Find(&cpuData).Error; err != nil {
			log.Println("Failed to query: ", err)
		}
		agentData.CPUs = append(agentData.CPUs, cpuData...)

		var ramData []models.RAM
		if err := db.Where("agent_id = ? AND timestamp >= ?", agent.ID, time30SecondsAgo).Find(&ramData).Error; err != nil {
			log.Println("Failed to query: ", err)
		}
		agentData.RAM = append(agentData.RAM, ramData...)

		var ssdData []models.SSD
		if err := db.Where("agent_id = ? AND timestamp >= ?", agent.ID, time30SecondsAgo).Find(&ssdData).Error; err != nil {
			log.Println("Failed to query: ", err)
		}
		agentData.SSD = append(agentData.SSD, ssdData...)

		data.Agents = append(data.Agents, agentData)
	}

	response := map[string]models.ClientPackage{"data": data}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
