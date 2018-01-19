package db

import (
	"log"

	"github.com/domtheporcupine/divvyup_api/models"
)

/*
	GroupBalances is a function that returns a list
	of users and their balances given a group id
*/
func GroupBalances(gid int64) []models.UserJSON {
	ret := []models.UserJSON{}
	members := GroupMembers(gid)
	for _, member := range members {
		nmember := models.UserJSON{}
		nmember.Username = UserName(member)
		nmember.Balance = Balance(member)
		ret = append(ret, nmember)
	}
	return ret
}

/*
	Balance is a function that returns a users
	balance given a userid
*/
func Balance(uid int64) float64 {
	rows, err := db.Query("select balance from balances where userid = ?", uid)
	// If for some reason there is an error
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var balance float64

		if err := rows.Scan(&balance); err != nil {
			// Something has gone wrong
			log.Fatal(err)
		}

		return balance
	}

	return -1
}
