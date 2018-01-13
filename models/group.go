package models

/*
	Group is the model we will use to create a new group when a user
	requests one
*/
type Group struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}
