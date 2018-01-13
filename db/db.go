package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/domtheporcupine/divvyup_api/config"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

/*
	Init will set up all the necessary things to connect to
	our database
*/
func Init() {
	os.Remove("./divvyup_db.db")
	var err error
	db, err = sql.Open("sqlite3", "./divvyup_db.db")

	if err != nil {
		panic(err)
	}

	// TODO, figure out why this closes immediately
	// defer db.Close()

	fmt.Printf("Initializing database...")
	sqlStmt, err := ioutil.ReadFile(config.SchemaFile())

	if err != nil {
		fmt.Printf("\t\t\tError reading schema file: '%s'\n", config.SchemaFile())
		os.Exit(2)
	}

	_, err = db.Exec(string(sqlStmt[:]))

	db.Begin()

	if err != nil {
		panic(err)
	}

	fmt.Println("\t\t\tDone!")
}

/*
	CreateUser is a function to create a user in the database
	given a username and a password

	Return true on success, failure otherwise
*/
func CreateUser(uName string, pass string) bool {
	// First let's hash the password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {
		return false
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return false
	}
	stmt, err := tx.Prepare("insert into users(username, password) values(?, ?)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(uName, string(hashPass[:]))
	if err != nil {
		log.Fatal(err)
		return false
	}
	tx.Commit()
	return true
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

	Returns true if credentials are valid, false otherwise
*/
func AuthenticateUser(uName string, pass string) bool {
	rows, err := db.Query("select password,COUNT(*) from users where username = ?", uName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows == nil {
		fmt.Println("foobar")
	}
	for rows.Next() {
		var count int
		var prevpass string

		if err := rows.Scan(&prevpass, &count); err != nil {
			return false
		}
		// There is at least 1 user
		if count != 0 {
			// Compare the passwords
			err := bcrypt.CompareHashAndPassword([]byte(prevpass), []byte(pass))

			if err == nil {
				// They mathch!
				return true
			}
		}
	}
	return false
}
