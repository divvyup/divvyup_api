package api

import (
	"net/http"

	"github.com/domtheporcupine/divvyup_api/config"
	"github.com/gorilla/mux"
)

type Middleware func(w http.ResponseWriter, r *http.Request)

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
	router = AddItemRoutes(subrouter)
	router = AddUserRoutes(subrouter)
	return router
}

func Preflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Request-Headers", "X-Requested-With")
	w.Header().Set("Access-Control-ALlow-Headers", "content-type")
	w.WriteHeader(200)
	return
}
