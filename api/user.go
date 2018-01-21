package api

import (
	"encoding/json"
	"net/http"

	"github.com/domtheporcupine/divvyup_api/models"

	"github.com/gorilla/mux"
)

/*
	AddUserRoutes adds the following functionality:
	1. get info about a user			/user					GET
	2. update a user 							/user/{id}		UPDATE
	3. delete a user							/user/{id}		DELETE
*/
func AddUserRoutes(router *mux.Router) *mux.Router {
	// Add all our routes and handler functions
	router.Path("/user").HandlerFunc(Validate(getUserInfoHandler)).Methods("GET")
	// Send the new router back
	return router
}

func getUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	usr, _ := req.Context().Value(models.User{}).(models.User)
	res, _ := json.Marshal(models.UserJSON{Username: usr.Username})
	w.Write(res)
	return
}
