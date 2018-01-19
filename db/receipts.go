package db

import (
	"log"
)

/*
	CreateReceipt is a function to create a receipt in the database
	given a groupid

	Return true on success, failure otherwise
*/
func CreateReceipt(gid int64) bool {
	// Sanity check on group id
	if !ValidID(gid) {
		return false
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return false
	}
	stmt, err := tx.Prepare("insert into receipts(groupid) values(?)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(gid)
	if err != nil {
		log.Fatal(err)
		return false
	}
	tx.Commit()
	return true
}

/*
	BelongToGroup is a function that returns a slice
	of receipt id's belonging to the specified group
*/
func BelongToGroup(gid int64) []int64 {
	var rids []int64
	// Even though it has probably already
	// been done check if the gid is valid
	if !ValidID(gid) {
		return rids
	}

	rows, err := db.Query("select id from receipts where groupid = ?", gid)
	// If for some reason there is an error
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var nid int64
		if err := rows.Scan(&nid); err != nil {
			rids = append(rids, nid)
		}
	}

	return rids
}