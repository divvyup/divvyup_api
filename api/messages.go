package api

/*
	Message is our generic response structure
	it allows us to convey useful information
	to the client
*/
type Message struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
