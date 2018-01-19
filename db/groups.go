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
	// Sanity check that the id is valid
	ValidID(gid)

	// The group might exist so try to delete it
	DeleteByID(DelTableStrings.Group, gid)
	// Now we have to get all of the receipts that
	// were associated with the account

	// In this case we don't want to kill our
	// entire program if the wrong gid is given

	return true
}

/*
	GroupName is a function that returns the name of
	a group given a group id

	returns the name on success, the empty string otherwise
*/
func GroupName(gid int64) string {
	rows, err := db.Query("select name from groups where id = ?", gid)
	// If for some reason there is an error
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			// Something has gone wrong
			log.Fatal(err)
		}

		return name
	}

	return ""
}

/*
	GroupMembers is a function that returns
	a slice of id's that belong to the members
	of the group
*/
func GroupMembers(gid int64) []int64 {
	ids := []int64{}
	rows, err := db.Query("select userid from membership where groupid = ?", gid)
	// If for some reason there is an error
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64

		if err := rows.Scan(&id); err != nil {
			// Something has gone wrong
			log.Fatal(err)
		}

		ids = append(ids, id)
	}

	return ids
}
