package apis

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Route is
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is
type Routes []Route

// NewRouter is
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Index is
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	/* Route{
		"Index",
		"GET",
		"/",
		Index,
	}, */

	Route{
		"PersonsGet",
		"GET",
		"/persons",
		PersonsGet,
	},

	Route{
		"SocketHandler",
		"GET",
		"/socket.io/",
		SocketHandler,
	},
}
