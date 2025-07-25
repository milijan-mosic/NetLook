package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	database "github.com/milijan-mosic/net-look/server/database"
	models "github.com/milijan-mosic/net-look/server/models"
)

func ReceiveMetrics(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var data models.Package
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusInternalServerError)
		return
	}

	db := database.GetDB()
	allAgents := database.FindRootAgent(db)
	if len(allAgents) == 0 {
		http.Error(w, "No agents found", http.StatusInternalServerError)
		return
	}

	agentID := allAgents[0].ID
	date := data.Time.Date
	timestamp := data.Time.Timestamp

	var CPUDataForInsertion []models.CPU
	for _, receivedCPUData := range data.Cpu {
		newCPU := models.CPU{
			ID:        uuid.New().String(),
			AgentID:   agentID,
			Number:    receivedCPUData.Number,
			Usage:     receivedCPUData.Usage,
			Date:      date,
			Timestamp: timestamp,
		}
		CPUDataForInsertion = append(CPUDataForInsertion, newCPU)
	}
	if err := database.CreateModels(db, CPUDataForInsertion, false); err != nil {
		log.Println("Failed to create models: ", err)
		http.Error(w, "Failed to create models", http.StatusInternalServerError)
		return
	}

	newRAM := models.RAM{
		ID:        uuid.New().String(),
		AgentID:   agentID,
		Usage:     data.Ram.Usage,
		Used:      data.Ram.Used,
		Total:     data.Ram.Total,
		Date:      date,
		Timestamp: timestamp,
	}
	if err := database.CreateModel(db, newRAM, false); err != nil {
		log.Println("Failed to create model: ", err)
		http.Error(w, "Failed to create model", http.StatusInternalServerError)
		return
	}

	newSSD := models.SSD{
		ID:        uuid.New().String(),
		AgentID:   agentID,
		Usage:     data.Ssd.Usage,
		Used:      data.Ssd.Used,
		Total:     data.Ssd.Total,
		Date:      date,
		Timestamp: timestamp,
	}
	if err := database.CreateModel(db, newSSD, false); err != nil {
		log.Println("Failed to create model: ", err)
		http.Error(w, "Failed to create model", http.StatusInternalServerError)
		return
	}

	log.Println("Received...")

	response := map[string]string{"message": "Package received"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
