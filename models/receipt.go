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

/*
	ReceiptJSON is our model
	for the information related
	to a single receipt
*/
type ReceiptJSON struct {
	Items []Item  `json:"items"`
	Total float64 `json:"total"`
}

/*
	ReceiptsJSON is our model
	for handling *MULTIPLE*
	receipts
*/
type ReceiptsJSON struct {
	Receipts []int64 `json:"receipts"`
}
