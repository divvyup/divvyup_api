package api

/*
	AddAuthRoutes is a function that will add all of the functionality of
	our authentication related routes to the app
*/import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// AddHelloRoutes does
func AddHelloRoutes(router *mux.Router) *mux.Router {

	router.Path("/hello").HandlerFunc(Validate(http.HandlerFunc(helloHandler))).Methods("POST")

	return router
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	foo, _ := r.Context().Value(UserInfo{}).(UserInfo)
	fmt.Println(foo.Userid)
	clms, ok := r.Context().Value(UserInfo{}).(UserInfo)
	if !ok {
		w.Write([]byte("Nope."))
		return
	}
	fmt.Println(clms.Userid)
	w.Write([]byte("Success."))
}
