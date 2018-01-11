package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/domtheporcupine/divvyup_api/models"
	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

/*
	AddAuthRoutes is a function that will add all of the functionality of
	our authentication related routes to the app
*/
func AddAuthRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/hello", helloHandler).Methods("GET")
	router.HandleFunc("/auth/login", loginHandler).Methods("POST")
	router.HandleFunc("/auth/register", registerHandler).Methods("POST")
	return router
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	eUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&eUser)

	for _, us := range AllUsers {
		if us.Username == eUser.Username {
			err := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(eUser.Password))
			if err == nil {
				w.Write([]byte("Success!"))
			} else {
				w.Write([]byte("Failure!"))
			}
			return
		}
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	nUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&nUser)

	for _, us := range AllUsers {
		if us.Username == nUser.Username {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(nUser.Password), 14)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	nUser.Password = string(bytes[:])

	// Since we are here we know there is no other
	// user with that username so add the new user
	// to the 'database'
	AllUsers = append(AllUsers, *nUser)
	printUsers()
	w.Write([]byte("Success!"))
}

func printUsers() {
	for _, us := range AllUsers {
		fmt.Printf("%s has a password: %s\n", color.BlueString(us.Username), color.RedString(us.Password))
	}
}
