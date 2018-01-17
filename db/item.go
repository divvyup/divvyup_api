package db

import "log"

/*
	CreateItem is a function that will allow us
	to insert a new item into the database given
	an item name and a receipt id

	return true on success, false otherwise
*/
func CreateItem(rid int64, name string) bool {
	// Make sure the rid is valid
	if !ValidID(rid) {
		return false
	}

	// TODO make sure that receiptid is valid

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return false
	}
	stmt, err := tx.Prepare("insert into items(receiptid, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(rid, name)
	if err != nil {
		log.Fatal(err)
		return false
	}
	tx.Commit()

	return true
}
