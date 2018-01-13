package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/domtheporcupine/divvyup_api/config"
	"github.com/domtheporcupine/divvyup_api/db"
	"github.com/domtheporcupine/divvyup_api/models"

	jwt "github.com/dgrijalva/jwt-go"
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

/*
	Login Handler
*/
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// We will be responding with json
	w.Header().Set("Content-Type", "application/json")

	// Parse out the login info
	eUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&eUser)
	if db.AuthenticateUser(eUser.Username, eUser.Password) {
		// Username and password matches, time to give them a token
		// Declare the token we will be giving them
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		// Set their userid
		claims["userid"] = eUser.Username
		// Make sure the token experies in a reasonable amount of time
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		tokenString, err := token.SignedString(config.Secret())
		if err != nil {
			log.Fatal(err)
		}
		// Create our authorization cookie with the new token
		cookie := http.Cookie{
			Name:     "Authorization",
			Value:    tokenString,
			Expires:  time.Now().Add(time.Hour * 2),
			HttpOnly: true,
			Path:     "/",
		}

		http.SetCookie(w, &cookie)

		res, _ := json.Marshal(Message{Message: "Login successful.", Reason: "success"})
		w.Write(res)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	res, _ := json.Marshal(Message{Message: "Login unsuccessful.", Reason: "failure"})
	w.Write(res)
}

/*
	Registration handler
*/
// TODO: add checks for password length
// TODO: add second password for confirmation
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// We will be responding with json
	w.Header().Set("Content-Type", "application/json")
	// Parse out the requested information
	nUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&nUser)

	// Check if username is taken and respond appropriately
	if db.UserExists(nUser.Username) {
		res, _ := json.Marshal(Message{Message: "That username is already taken.", Reason: "username_taken"})
		w.Write(res)
		return
	}

	if db.CreateUser(nUser.Username, nUser.Password) {
		res, _ := json.Marshal(Message{Message: "Successfully created user.", Reason: "success"})
		w.Write(res)
	} else {
		res, _ := json.Marshal(Message{Message: "User could not be created.", Reason: "failure"})
		w.Write(res)
	}
}
