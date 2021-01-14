package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/peterjochum/traffic-jam-api/pkg/store"
)

const (
	devMode        = "dev"
	prodMode       = "prod"
	modeEnvVarname = "TJ_MODE"
	portEnvVarname = "TJ_PORT"
	defaultPort    = 8080
)

// GetServerPort reads  the port from environment or assumes 8080 as default
func GetServerPort() (int, error) {
	portFromEnv := os.Getenv(portEnvVarname)
	if portFromEnv != "" {
		var err interface{}
		port, err := strconv.Atoi(portFromEnv)
		if err != nil {
			msg := fmt.Sprintf("%s environment variable should be a port number", portEnvVarname)
			return 0, errors.New(msg)
		}
		return port, nil
	}
	return defaultPort, nil
}

// SetupStores initializes data stores based on the selected environment
func SetupStores() error {
	var tjs store.TrafficJamStore
	mode := os.Getenv(modeEnvVarname)
	if mode == "" {
		mode = devMode
	}
	if mode == devMode {
		// Use in memory store for the time being
		tjs = store.NewInMemoryTrafficJamStore()
		store.SeedTrafficJamStore(tjs)
	} else {
		return fmt.Errorf("Currently TJ_MODE only supports %s", devMode)
	}
	log.Printf("Using environment %s ..", mode)
	GlobalTrafficJamStore = tjs
	return nil
}
