package main

import (
	"log"
	"net/http"

	controllers "github.com/milijan-mosic/net-look/server/controllers"
	database "github.com/milijan-mosic/net-look/server/database"
	models "github.com/milijan-mosic/net-look/server/models"
)

func main() {
	db := database.OpenDBConnection("./server.db")
	database.MigrateModels(db, models.AllModels)
	database.SeedDatabase(db)
	database.CloseDBConnection(db, false)

	http.HandleFunc("/receive", controllers.ReceiveMetrics)
	http.HandleFunc("/test", controllers.Test)
	log.Fatal(http.ListenAndServe(":11000", nil))
}
