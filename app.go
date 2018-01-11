package main

import (
	"net/http"

	"github.com/domtheporcupine/divvyup_api/api"
	"github.com/domtheporcupine/divvyup_api/config"
)

func main() {
	config.AppConfig()
	router := api.InitRoutes()
	http.ListenAndServe(":3030", router)
}
