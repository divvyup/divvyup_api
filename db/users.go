package db

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

/*
	CreateUser is a function to create a user in the database
	given a username and a password

	Return new user id on success, -1 otherwise
*/
// TODO: checks on password length/strength
func CreateUser(uName string, pass string) int64 {
	// First let's hash the password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {
		return -1
	}

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into users(username, password) values(?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(uName, string(hashPass[:]))
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
	UserExists is a function to determine if a given user exists
	it is useful in a number of places so abstracting it our makes
	sense
*/
func UserExists(uName string) bool {
	rows, err := db.Query("select COUNT(*) from users where username = ?", uName)
	// If for some reason there is an error
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var count int

		if err := rows.Scan(&count); err != nil {
			// Something has gone wrong
			log.Fatal(err)
		}

		if count != 0 {
			return true
		}
	}

	return false
}

/*
	AuthenticateUser is a function to check if a user login
	is valid or not

	Returns the user's id if credentials are valid, -1 otherwise
*/
func AuthenticateUser(uName string, pass string) int64 {
	rows, err := db.Query("select id,password,COUNT(*) from users where username = ?", uName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows == nil {
		fmt.Println("foobar")
	}
	for rows.Next() {
		var id int64
		var count int
		var prevpass string

		if err := rows.Scan(&id, &prevpass, &count); err != nil {
			return -1
		}
		// There is at least 1 user
		if count != 0 {
			// Compare the passwords
			err := bcrypt.CompareHashAndPassword([]byte(prevpass), []byte(pass))

			if err == nil {
				// They mathch!
				return id
			}
		}
	}
	return -1
}

/*
	IsMember is a function to determine
	if a given user is a member of a specified
	group

	return true if the user is a member, false otherwise
*/
func IsMember(uid int64, gid int64) bool {
	// First validate the ids
	if !ValidID(uid) || !ValidID(gid) {
		return false
	}

	rows, err := db.Query("select COUNT(*) from membership where userid = ? and groupid =?", uid, gid)
	// If for some reason there is an error
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var count int

		if err := rows.Scan(&count); err != nil {
			// Something has gone wrong
			log.Fatal(err)
		}

		if count != 0 {
			return true
		}
	}

	return false
}

/*
	UserName is a function to get a users username
	given a user id

	returns the users username on success, the
	empty string otherwise
*/
func UserName(uid int64) string {
	rows, err := db.Query("select username from users where id = ?", uid)
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
