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

/*
	BelongToReceipt is a function that returns a slice
	of item id's belonging to the specified receipt
*/
func BelongToReceipt(rid int64) []int64 {
	var iids []int64
	// Even though it has probably already
	// been done check if the gid is valid
	if !ValidID(rid) {
		return iids
	}

	rows, err := db.Query("select id from items where receiptid = ?", rid)
	// If for some reason there is an error
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var nid int64
		if err := rows.Scan(&nid); err != nil {
			iids = append(iids, nid)
		}
	}

	return iids
}
