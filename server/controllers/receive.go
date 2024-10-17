package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

	fmt.Println(data)
}
