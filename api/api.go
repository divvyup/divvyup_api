package api

import (
	"github.com/domtheporcupine/divvyup_api/config"
	"github.com/gorilla/mux"
)

/*
	InitRoutes is the backbone of our api, it organizes everything
	and ties everything together at the highest level. It returns
	our router to the main part of our app to use
*/
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	// Here we use our config to define what our api version should be
	subrouter := router.PathPrefix(config.Prefix()).Subrouter()
	router = AddAuthRoutes(subrouter)
	router = AddGroupRoutes(subrouter)
	return router
}
