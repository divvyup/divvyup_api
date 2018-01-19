package models

/*
	Group is the model we will use to create a new group when a user
	requests one
*/
type Group struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

/*
	GroupJSON is the model we use to respond to a request about a
	group, it gives useful information for the fron end
*/
type GroupJSON struct {
	Name    string     `json:"name"`
	ID      int64      `json:"id"`
	Members []UserJSON `json:"members"`
}
