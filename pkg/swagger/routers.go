package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Route stores http and handler information of an API route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a list of all supported routes
type Routes []Route

// NewRouter returns the API routing with middlewares attached
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

// Index contains a welcome page that leads to the API endpoints
func Index(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "/api/v1/trafficjam")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"Index",
		"GET",
		"/api/v1/",
		Index,
	},

	Route{
		"AddTrafficJam",
		strings.ToUpper("Post"),
		"/api/v1/trafficjam",
		AddTrafficJam,
	},

	Route{
		"DeleteTrafficJam",
		strings.ToUpper("Delete"),
		"/api/v1/trafficjam/{id}",
		DeleteTrafficJam,
	},

	Route{
		"GetAllTrafficJams",
		strings.ToUpper("Get"),
		"/api/v1/trafficjam",
		GetAllTrafficJams,
	},

	Route{
		"GetTrafficJam",
		strings.ToUpper("Get"),
		"/api/v1/trafficjam/{id}",
		GetTrafficJam,
	},

	Route{
		"PutTrafficJam",
		strings.ToUpper("Put"),
		"/api/v1/trafficjam",
		PutTrafficJam,
	},
}
