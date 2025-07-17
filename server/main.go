package main

import (
	"log"
	"net/http"

	controllers "github.com/milijan-mosic/net-look/server/controllers"
	database "github.com/milijan-mosic/net-look/server/database"
	models "github.com/milijan-mosic/net-look/server/models"
)

var url = "/api/1.0"

func initDatabase() {
	db := database.OpenDBConnection("./server.db")
	database.MigrateModels(db, models.AllModels)
	database.SeedDatabase(db)
	database.CloseDBConnection(db, false)
}

func main() {
	initDatabase()

	http.HandleFunc(url+"/", controllers.ServeClient)
	http.HandleFunc(url+"/receive", controllers.ReceiveMetrics)
	http.HandleFunc(url+"/test", controllers.Test)

	log.Fatal(http.ListenAndServe(":10000", http.DefaultServeMux))
}
