package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/domtheporcupine/divvyup_api/config"
	_ "github.com/mattn/go-sqlite3"
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

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		// do something here
		fmt.Println("bad things")
	}

	fmt.Printf("Initializing database...")
	sqlStmt, err := ioutil.ReadFile(config.SchemaFile())

	if err != nil {
		fmt.Printf("Error reading schema file: '%s'\n", config.SchemaFile())
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

	_, err = stmt.Exec(uName, pass)
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
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	defer rows.Close()

	if rows == nil {
		return false
	}

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
