package main

import (
	"log"
	"net/http"

	controllers "github.com/milijan-mosic/net-look/server/controllers"
)

func main() {
	http.HandleFunc("/receive", controllers.ReceiveMetrics)
	log.Fatal(http.ListenAndServe(":11000", nil))
}
