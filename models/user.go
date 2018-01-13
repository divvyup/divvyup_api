package models

/*
	User is the model we will use to contain all the different attriubtes of
	well, a user.
*/
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
}
