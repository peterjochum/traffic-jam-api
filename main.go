// package main contains the initialization logic of the application
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/peterjochum/traffic-jam-api/pkg/app"
	"github.com/peterjochum/traffic-jam-api/pkg/store"
	sw "github.com/peterjochum/traffic-jam-api/pkg/swagger"
)

const (
	devMode        = "dev"
	prodMode       = "prod"
	portEnvVarname = "TJ_PORT"
)

func main() {

	setupStores()
	port := getServerPort()
	router := sw.NewRouter()
	log.Printf("Starting TrafficJam server on :%d", port)
	listenAddress := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(listenAddress, router))
}

func getServerPort() int {
	portFromEnv := os.Getenv(portEnvVarname)
	if portFromEnv != "" {
		var err interface{}
		port, err := strconv.Atoi(portFromEnv)
		if err != nil {
			log.Fatalf("%s environment variable should be a port number", portEnvVarname)
		}
		return port
	} else {
		return 8080
	}
}

func setupStores() {
	var tjs store.TrafficJamStore
	mode := os.Getenv("TJ_MODE")
	if mode == "" {
		mode = devMode
	}
	if mode == devMode {
		// Use in memory store for the time being
		tjs = store.NewInMemoryTrafficJamStore(true)
	} else {
		log.Fatalf("Currently TJ_MODE only supports %s", devMode)
	}
	log.Printf("Using environment %s ..", mode)
	app.TrafficJamStore = tjs
}
