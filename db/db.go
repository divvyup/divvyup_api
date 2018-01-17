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

/*
	TableStrings is a struct we will use to hold information
	about all the tables we have, it is useful in generalizing
	the process of inseting and deleting from tables using
	things like ids
*/
type TableStrings struct {
	User       string
	Membership string
	Group      string
	Receipt    string
	Item       string
}

var db *sql.DB
var delTableStrings *TableStrings

/*
	Init will set up all the necessary things to connect to
	our database
*/
func Init() {
	if config.DBDriver() == "sqlite3" {
		os.Remove(config.DBUrl())
	}

	var err error
	db, err = sql.Open(config.DBDriver(), config.DBUrl())

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
	delTableStrings = new(TableStrings)

	// Initialize all of our new strings
	delTableStrings.User = "delete from users where id = ?"
	delTableStrings.Membership = "delete from membership where id = ?"
	delTableStrings.Group = "delete from groups where id = ?"
	delTableStrings.Receipt = "delete from receipts where id = ?"
	delTableStrings.Item = "delete from items where id = ?s"

	fmt.Println("\t\t\tDone!")
}

/*
	CheckErr is a useful function. It eliminates
	repetative code, because we commonly check errors
	the same way
*/
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*
	ValidID is a function that will allow us to
	check if a given id meets a minimum set of requirments
	usually being positive
*/
func ValidID(id int64) bool {
	if id > 0 {
		return true
	}
	return false
}
