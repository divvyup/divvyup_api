package db

import (
	"math/rand"

	"github.com/domtheporcupine/divvyup_api/examples"
)

// This is starting to look and more like a test
// file but we are going to leave it for now

var usernames = []string{"dom", "brandon", "kristiana", "michaela", "me"}
var passwords = []string{"bestpassword", "secure_password", "even_more_secure_password", "security_expert", "password"}
var userids = []int64{}

var groups = []string{"roomies", "friends"}
var groupids = []int64{}

/*
	Populate is a function that we use in the demo
	to fill the database with some example receipts
	groups and items.
*/
func Populate() {
	// First let's make 5 user accounts
	for index, username := range usernames {
		userids = append(userids, CreateUser(username, passwords[index]))
	}

	// Now lets go ahead and make 2 groups
	for _, group := range groups {
		groupids = append(groupids, CreateGroup(group))
	}

	// Add some random users to the groups

	// Three into roomies
	roomies := rand.Perm(4)

	AddUserToGroup(userids[roomies[0]], groupids[0])
	AddUserToGroup(userids[roomies[1]], groupids[0])
	// Add the example user to the group
	AddUserToGroup(userids[4], groupids[0])

	// Add everyone to the group of friends
	for i := 0; i < 5; i++ {
		AddUserToGroup(userids[i], groupids[1])
	}

	// Create a fake receipt
	receiptid := CreateReceipt(groupids[0])

	// Add some sample items to the receipt
	items := examples.GetExampleItems()
	for _, item := range items {
		CreateItem(receiptid, item.Name, item.Price)
	}

}
