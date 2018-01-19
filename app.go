package main

import (
	"net/http"

	"github.com/domtheporcupine/divvyup_api/api"
	"github.com/domtheporcupine/divvyup_api/config"
	"github.com/domtheporcupine/divvyup_api/db"
)

func main() {
	config.AppConfig()
	db.Init()
	db.Populate()
	router := api.InitRoutes()
	http.ListenAndServe(":3030", router)
}
