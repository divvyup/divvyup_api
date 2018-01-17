package db

import (
	"log"
)

/*
	CreateGroup is a function to create a group in the database
	given a name

	Return the new id on success -1 otherwise
*/
func CreateGroup(name string) int64 {

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return -1
	}

	// Set up for our new group
	stmt, err := tx.Prepare("insert into groups(name) values(?)")
	if err != nil {
		log.Fatal(err)
		return -1
	}
	defer stmt.Close()

	res, err := stmt.Exec(name)

	if err != nil {
		log.Fatal(err)
		return -1
	}
	tx.Commit()

	// Finally grab the new group id!
	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
		return -1
	}
	return id
}

/*
	AddUserToGroup is a function to add a new user to a group

	return true on success, false otherwise
*/
func AddUserToGroup(uid int64, gid int64) bool {
	tx, err := db.Begin()
	CheckErr(err)

	// Set up for our new group
	stmt, err := tx.Prepare("insert into membership(userid,groupid) values(?,?)")
	CheckErr(err)

	defer stmt.Close()

	_, err = stmt.Exec(uid, gid)

	CheckErr(err)

	tx.Commit()
	return true
}

/*
	DeleteGroup is a function to delete and existing group

	It is slightly more complicated than a simple delete because
	we must go down the chain and delete *EVERYTHING* associated
	with the account, this includes all receipts, and all those
	receipts items.

	return true on success, false otherwise
*/
func DeleteGroup(gid int64) bool {
	// Sanity check that the number is
	// greater than 0
	if gid < 1 {
		return false
	}

	// The group might exist so try to delete it
	stmt, err := db.Prepare("delete from groups where id=?")
	CheckErr(err)

	res, err := stmt.Exec(gid)

	// Now we have to get all of the receipts that
	// were associated with the account

	// In this case we don't want to kill our
	// entire program if the wrong gid is given

	_, err = res.RowsAffected()
	CheckErr(err)

	return true
}
