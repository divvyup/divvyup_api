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
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Set up for our new group
	stmt, err := tx.Prepare("insert into membership(userid,groupid) values(?,?)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(uid, gid)

	if err != nil {
		log.Fatal(err)
		return false
	}
	tx.Commit()
	return true
}
