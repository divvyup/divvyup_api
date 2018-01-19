package main

import (
	"net/http"

	"github.com/domtheporcupine/divvyup_api/api"
	"github.com/domtheporcupine/divvyup_api/config"
	"github.com/domtheporcupine/divvyup_api/db"
	"github.com/domtheporcupine/divvyup_api/examples"
)

func main() {
	config.AppConfig()
	db.Init()
	examples.ParseReceiptData()
	db.Populate()
	router := api.InitRoutes()
	http.ListenAndServe(":3030", router)
}
