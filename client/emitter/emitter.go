package emitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "github.com/milijan-mosic/net-look/client/models"
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
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if debug {
		fmt.Println("Package status:", resp.Status)
	}
}
