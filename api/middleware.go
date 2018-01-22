package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/domtheporcupine/divvyup_api/db"
	"github.com/gorilla/mux"

	"github.com/domtheporcupine/divvyup_api/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/domtheporcupine/divvyup_api/config"
)

/*
	Validate is the main middleware function. We use it in between all
	routes that need to be authenticated and their handler functions
*/
func Validate(page Middleware) http.HandlerFunc {
	protectedPage := http.HandlerFunc(page)
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "OPTIONS" {
			protectedPage = http.HandlerFunc(Preflight)
			protectedPage.ServeHTTP(w, req)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", config.Host())
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Request-Headers", "X-Requested-With")
		// If no Auth cookie is set then return a 404 not found
		cookie, err := req.Cookie("Authorization")
		if err != nil {
			http.NotFound(w, req)
			return
		}
		user, err := authorize(cookie)

		if err == nil {
			ctx := context.WithValue(req.Context(), models.User{}, *user)
			protectedPage.ServeHTTP(w, req.WithContext(ctx))
		} else {
			http.NotFound(w, req)
			return
		}

	})
}

/*
	ValidateWithGroup is a middleware function to validate
	whether a user belongs to a group, as this is a common
	thing to do
*/
func ValidateWithGroup(page Middleware) http.HandlerFunc {
	protectedPage := http.HandlerFunc(page)
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "OPTIONS" {
			protectedPage = http.HandlerFunc(Preflight)
			protectedPage.ServeHTTP(w, req)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", config.Host())
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Request-Headers", "X-Requested-With")
		// Grab our auth cookie
		cookie, err := req.Cookie("Authorization")

		if err != nil {
			http.NotFound(w, req)
			return
		}

		user, err := authorize(cookie)

		var gid int64
		gid = -1

		if req.Method == "POST" {
			// See if we can pull out the group information
			nGroup := new(models.Group)
			json.NewDecoder(req.Body).Decode(&nGroup)
			gid = nGroup.ID
		} else if req.Method == "GET" {
			vars := mux.Vars(req)
			gid, err = strconv.ParseInt(vars["id"], 10, 64)
			if err != nil {
				gid = -1
			}
		}

		if db.IsMember(user.ID, gid) {
			ctx := context.WithValue(req.Context(), models.User{}, *user)
			protectedPage.ServeHTTP(w, req.WithContext(ctx))
			return
		}

		res, _ := json.Marshal(Message{Message: "You do not belong to a group with that id.", Reason: "invalid_group_id"})
		w.Write(res)
		return
	})
}

/*
	authorize is the function that does all the heavy lifting, it allows
	us to have many different validation middleware functions

	given a cookie it validates, pulls out the claims and returns
	a user model
*/

func authorize(cookie *http.Cookie) (*models.User, error) {
	// Return a Token using the cookie
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's signature wasn't changed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return config.Secret(), nil
	})

	if err != nil {
		return nil, err
	}

	// Grab the tokens claims and pass it into the original request
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		currUser := new(models.User)
		currUser.Username = claims["username"].(string)
		currUser.ID = int64(claims["id"].(float64))
		fmt.Println(currUser.Username)
		return currUser, nil
	}
	return nil, fmt.Errorf("error validating user")
}

func CORS(page Middleware) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// We will be responding with json
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", config.Host())
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Request-Headers", "X-Requested-With")
		npage := http.HandlerFunc(page)
		if req.Method == "OPTIONS" {
			npage = http.HandlerFunc(Preflight)
		}
		npage.ServeHTTP(w, req)
	})
}
