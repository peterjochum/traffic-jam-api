// package main contains the initialization logic of the application
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/peterjochum/traffic-jam-api/internal/app"
	sw "github.com/peterjochum/traffic-jam-api/pkg/swagger"
)

// main wires together components and launches the server
func main() {
	// Initialize stores
	if err := app.SetupStores(); err != nil {
		log.Fatal(err)
	}

	// Get server parameters
	port, err := app.GetServerPort()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize routing
	router := sw.NewRouter()

	// Start the server
	log.Printf("Starting TrafficJam server on :%d", port)
	listenAddress := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(listenAddress, router))
}
