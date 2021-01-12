// package main contains the initialization logic of the application
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/peterjochum/traffic-jam-api/pkg/app"
	"github.com/peterjochum/traffic-jam-api/pkg/store"
	sw "github.com/peterjochum/traffic-jam-api/pkg/swagger"
)

const (
	devMode  = "dev"
	prodMode = "prod"
)

func main() {
	log.Printf("Starting TrafficJam server")

	var tjs store.TrafficJamStore
	mode := os.Getenv("TJ_MODE")
	if mode == devMode {
		// Use in memory store for the time being
		tjs = store.NewInMemoryTrafficJamStore(true)
	} else {
		log.Fatalf("Set environment variable TJ_MODE to %s/%s", devMode, prodMode)
	}
	log.Printf("Using environment %s ..", mode)
	app.TrafficJamStore = tjs

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
