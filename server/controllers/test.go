package controllers

import (
	"log"
	"net/http"

	database "github.com/milijan-mosic/net-look/server/database"
)

func Test(w http.ResponseWriter, r *http.Request) {
	url := database.GetDBUrl()
	db := database.OpenDBConnection(url)

	var tables []struct {
		Name string `gorm:"column:name"`
	}

	if err := db.Raw("SELECT name FROM sqlite_master WHERE type='table';").Scan(&tables).Error; err != nil {
		log.Fatal("Failed to fetch table names: ", err)
	}

	database.CloseDBConnection(db, false)

	for _, table := range tables {
		log.Println("Table Name: ", table.Name)
	}
}
