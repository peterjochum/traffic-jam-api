package handlers

import (
	"fmt"
	"net/http"
)

const apiHelp = "/api/v1/trafficjam"

// Index contains a welcome page that leads to the API endpoints
func Index(w http.ResponseWriter, _ *http.Request) {

	_, _ = fmt.Fprint(w, apiHelp)
}
