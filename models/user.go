package models

/*
	User is the model we will use to contain all the different attriubtes of
	well, a user.
*/
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int64  `json:"id"`
}

/*
	UserJSON is the model we will use for sending user information
	to the clients. We do not want to be sending the password and
	id, but often will want to spend their balance
*/
type UserJSON struct {
	Username string  `json:"name"`
	Balance  float64 `json:"balance,omitempty"`
	Groups   []int64 `json:"groups,omitempty"`
}
