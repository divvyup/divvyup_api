package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/domtheporcupine/divvyup_api/db"

	"github.com/domtheporcupine/divvyup_api/models"
	"github.com/gorilla/mux"
)

/*
	AddGroupRoutes adds the following functionality:
	1. create a new group						/gruop				POST
	2. add a user to a group			 	/group/{id}		UPDATE
	3. remove a user from a group		/group/{id}		UPDATE
	4. delete a group								/group/{id}		DELETE
	5. get info about a gorup				/group/{id}		GET
*/
func AddGroupRoutes(router *mux.Router) *mux.Router {

	router.Path("/group").HandlerFunc(Validate(http.HandlerFunc(createGroupHandler))).Methods("POST")
	router.Path("/group/{id}").HandlerFunc(Validate(http.HandlerFunc(getGroupInfoHandler))).Methods("GET")
	return router
}

func createGroupHandler(w http.ResponseWriter, r *http.Request) {
	// We will be responding with json
	w.Header().Set("Content-Type", "application/json")

	// First we need the current users info
	usr, _ := r.Context().Value(models.User{}).(models.User)

	// Next we need to parse out our new group's info
	nGroup := new(models.Group)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&nGroup)

	// Make our new group!
	gid := db.CreateGroup(nGroup.Name)

	// Check to see if we were succeful
	if gid != -1 {
		// If we were succeful we need to add ourselves to the group
		if db.AddUserToGroup(usr.ID, gid) {
			res, _ := json.Marshal(Message{Message: "Group created successfully.", Reason: "success"})
			w.Write(res)
			return
		}

		// If we could not add the user to the group
		// delete the group and then return an error
		db.DeleteGroup(gid)

	}
	w.WriteHeader(http.StatusAccepted)
	res, _ := json.Marshal(Message{Message: "Failed to create group.", Reason: "failure"})
	w.Write(res)

}

/*
	Given a valid group id this route responds with a json
	object in the form:

	{
		"name": "group_name"
		"id": "group_id"
		"members": [
			{"name": "user_name", "balance": "-42"},
			{"name": "user_2_name", "balance": "42"}
		]
	}

*/

func getGroupInfoHandler(w http.ResponseWriter, r *http.Request) {
	// We will be responding with json
	w.Header().Set("Content-Type", "application/json")
	// Pull out the group id
	vars := mux.Vars(r)
	groupID, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		// We couldn't parse out the id
		res, _ := json.Marshal(Message{Message: "Invalid group ID.", Reason: "invalid_id"})
		w.Write(res)
		return
	}
	// Get the current users id
	usr, _ := r.Context().Value(models.User{}).(models.User)

	// Validate that the user is a member of the group
	if !db.IsMember(usr.ID, groupID) {
		res, _ := json.Marshal(Message{Message: "You do not belong to a group with that ID.", Reason: "invalid_id"})
		w.Write(res)
		return
	}

	// By now we know that the user is a member of the group
	// so actually pull out the important information and send
	// it back to the user

	resp := new(models.GroupJSON)
	resp.Name = db.GroupName(groupID)
	resp.ID = groupID
	resp.Members = db.GroupBalances(groupID)

	res, _ := json.Marshal(resp)
	w.Write(res)
	return
}
