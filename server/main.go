package main

import (
	"log"
	"net/http"

	controllers "github.com/milijan-mosic/net-look/server/controllers"
	database "github.com/milijan-mosic/net-look/server/database"
)

var url = "/api/1.0"

func main() {
	database.InitializeDatabase()

	http.HandleFunc(url+"/", controllers.ServeClient)
	http.HandleFunc(url+"/receive", controllers.ReceiveMetrics)
	http.HandleFunc(url+"/test", controllers.Test)

	log.Fatal(http.ListenAndServe(":10000", http.DefaultServeMux))
}
