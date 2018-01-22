package db

import (
	"log"

	"github.com/domtheporcupine/divvyup_api/models"
)

/*
	CreateReceipt is a function to create a receipt in the database
	given a groupid

	Return true on success, failure otherwise
*/
func CreateReceipt(gid int64) int64 {
	// Sanity check on group id
	if !ValidID(gid) {
		return -1
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into receipts(groupid) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(gid)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	// Finally grab the new user id!
	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	return id
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
		if err := rows.Scan(&nid); err == nil {
			rids = append(rids, nid)
		}
	}

	return rids
}

/*
	BelongToUser is a function that determines if
	a given user can view a receipt
*/
func BelongToUser(uid int64, rid int64) bool {
	if !ValidID(rid) || !ValidID(uid) {
		return false
	}

	rows, err := db.Query("select COUNT(*) from groups inner join receipts on receipts.groupid = groups.id inner join membership on groups.id = membership.groupid inner join users on users.id = membership.userid where users.id = ? and receipts.id = ?", uid, rid)
	// If for some reason there is an error
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var count int64
		if err := rows.Scan(&count); err != nil {
			if count == 1 {
				return true
			}
		}
	}

	return false

}

/*
	GetReceipt is a function that simply returns
	the important information about a receipt
*/
func GetReceipt(rid int64) models.ReceiptJSON {
	var receipt = models.ReceiptJSON{}
	// Sanity check
	if !ValidID(rid) {
		return receipt
	}
	// First we need to get all items that belong to
	// the receipt
	rows, err := db.Query("select id,name,price from items where receiptid = ?", rid)

	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var item = models.Item{}
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err == nil {
			receipt.Items = append(receipt.Items, item)
		}
	}

	if err != nil {
		return receipt
	}

	return receipt
}
