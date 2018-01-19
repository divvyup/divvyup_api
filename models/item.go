package models

/*
	Item is our model for the items that
	get added to the receipt, the idea of
	an item is that it is very general. i.e.
	an item can be monthly rent
*/
type Item struct {
	Name  string  `json:"name"`
	ID    int64   `json:"id"`
	Price float64 `json:"price"`
}
