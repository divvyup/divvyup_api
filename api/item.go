package api

import (
	"github.com/gorilla/mux"
)

/*
	AddItemRoutes adds the following functionality:
	1. create a new item						/item					POST
	2. update an item 							/item/{id}		UPDATE
	3. delete an item								/item/{id}		DELETE
	4. get info about an item				/item/{id}		GET
*/
func AddItemRoutes(router *mux.Router) *mux.Router {
	return router
}
