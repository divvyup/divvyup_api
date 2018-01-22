package api

import (
	"encoding/json"
	"net/http"

	"github.com/domtheporcupine/divvyup_api/db"

	"github.com/domtheporcupine/divvyup_api/models"
	"github.com/gorilla/mux"
)

/*
	AddReceiptRoutes adds the following functionality:
	1. create a new receipt						/receipt				POST
	2. update a receipt 							/receipt/{id}		UPDATE
	4. delete a receipt								/receipt/{id}		DELETE
	5. get info about an receipt			/receipt/{id}		GET
*/
func AddReceiptRoutes(router *mux.Router) *mux.Router {
	// Add all our routes and handler functions
	router.Path("/receipt").HandlerFunc(ValidateWithGroup(http.HandlerFunc(createReceiptHandler))).Methods("POST")
	router.Path("/receipt/{id}").HandlerFunc(Validate(getReceiptInfoHandler)).Methods("GET")

	// Send the new router back
	return router
}

/*
	To create a receipt we should recieve a request in the following
	form [JSON]:
	{
		id: 8675309
		date: 01/20/2018 -- this is not implemented yet
	}

	The server will respond with a Message if there is an error or
	a Receipt with the new id on success
*/
func createReceiptHandler(w http.ResponseWriter, r *http.Request) {
	// We will be responding with json
	w.Header().Set("Content-Type", "application/json")

	// First we need to parse out our new receipts's info
	nGroup := new(models.Group)
	json.NewDecoder(r.Body).Decode(&nGroup)

	nID := db.CreateReceipt(nGroup.ID)

	if nID == -1 {
		res, _ := json.Marshal(Message{Message: "Error creating receipt.", Reason: "internal_error"})
		w.Write(res)
		return
	}

	res, _ := json.Marshal(models.Receipt{ID: nID})
	w.Write(res)
	return
}

func getReceiptInfoHandler(w http.ResponseWriter, r *http.Request) {
	// We will be responding with json
	w.Header().Set("Content-Type", "application/json")
}
