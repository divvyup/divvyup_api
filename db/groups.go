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
