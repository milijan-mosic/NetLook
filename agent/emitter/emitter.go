package emitter

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	models "github.com/milijan-mosic/net-look/agent/models"
)

func PackJson(data models.Package) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	return jsonData
}

func Emit(jsonData []byte, target string, debug bool) {
	log.Println("Sending...")

	req, err := http.NewRequest("POST", target, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	postman := &http.Client{}
	resp, err := postman.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	if debug {
		log.Println("Package status: ", resp.Status)
	}
}
