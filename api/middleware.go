package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/domtheporcupine/divvyup_api/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/domtheporcupine/divvyup_api/config"
)

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

			currUser := new(models.User)
			currUser.Username = claims["username"].(string)

			currUser.ID = int64(claims["id"].(float64))
			ctx := context.WithValue(req.Context(), models.User{}, *currUser)
			protectedPage.ServeHTTP(w, req.WithContext(ctx))
		} else {

			http.NotFound(w, req)
			return
		}

	})
}
