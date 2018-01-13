package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/domtheporcupine/divvyup_api/db"
	"github.com/domtheporcupine/divvyup_api/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

/*
	AddAuthRoutes is a function that will add all of the functionality of
	our authentication related routes to the app
*/
func AddAuthRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/auth/login", loginHandler).Methods("POST")
	router.HandleFunc("/auth/register", registerHandler).Methods("POST")
	return router
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

	if db.UserExists(nUser.Username) {
		fmt.Println("The user exists!")
		w.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(Message{Message: "That username is already taken.", Reason: "username_taken"})
		w.Write(res)
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(nUser.Password), 14)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	nUser.Password = string(bytes[:])
	if db.CreateUser(nUser.Username, nUser.Password) {
		w.Write([]byte("Success!"))
	} else {
		w.Write([]byte("Failure!"))
	}
}
