package models

/*
	Receipt is our model for
	receipt information, useful
	for creating a new receipt
	etc
*/
type Receipt struct {
	ID   int64  `json:"id"`
	Date string `json:"date"`
}
