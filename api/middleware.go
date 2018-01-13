package api

import (
	"context"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/domtheporcupine/divvyup_api/config"
)

/*
	UserInfo is the conatiner we use to pass context information
	down the line after it has been parsed out of the cookie
*/
type UserInfo struct {
	Userid string
}

/*
	Validate is the main middleware function. We use it in between all
	routes that need to be authenticated and their handler functions
*/
func Validate(protectedPage http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// If no Auth cookie is set then return a 404 not found
		cookie, err := req.Cookie("Authorization")
		if err != nil {
			fmt.Println("Error with cookie")
			http.NotFound(w, req)
			return
		}

		// Return a Token using the cookie
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			// Make sure token's signature wasn't changed
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return config.Secret(), nil
		})

		if err != nil {

			http.NotFound(w, req)
			return
		}

		// Grab the tokens claims and pass it into the original request
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			foo := new(UserInfo)
			foo.Userid = claims["userid"].(string)
			ctx := context.WithValue(req.Context(), UserInfo{}, *foo)
			protectedPage.ServeHTTP(w, req.WithContext(ctx))
		} else {

			http.NotFound(w, req)
			return
		}

	})
}
