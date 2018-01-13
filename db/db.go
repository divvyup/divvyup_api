package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
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

	fmt.Println("\t\t\tDone!")
}
