package app

import (
	"os"
	"strconv"
	"testing"

	"github.com/peterjochum/traffic-jam-api/pkg/store"
)

func TestGetServerPort(t *testing.T) {
	port, err := GetServerPort()
	if err != nil {
		t.Errorf("no error expected, got: %v", err)
	}
	if port != defaultPort {
		t.Errorf("Expected port to be %d by default", defaultPort)
	}
}

func TestGetServerPort_Enverr(t *testing.T) {
	_ = os.Setenv(portEnvVarname, "abc")
	defer func() {
		_ = os.Unsetenv(portEnvVarname)
	}()
	_, err := GetServerPort()
	if err == nil {
		t.Errorf("expected an error")
	}
}

func TestGetServerPort_Env(t *testing.T) {
	expectedPort := 9090
	_ = os.Setenv(portEnvVarname, strconv.Itoa(expectedPort))
	defer func() {
		_ = os.Unsetenv(portEnvVarname)
	}()
	port, err := GetServerPort()
	if err != nil {
		t.Error(err)
	}

	if port != expectedPort {
		t.Errorf("expected port %d, got %d", expectedPort, port)
	}
}

func TestSetupStores_Unset(t *testing.T) {
	if err := SetupStores(); err != nil {
		t.Error(err)
	}
	_, ok := GlobalTrafficJamStore.(*store.InMemoryTrafficJamStore)
	if !ok {
		t.Errorf("expected an in memory store pointer")
	}
}

func TestSetupStores_Dev(t *testing.T) {
	_ = os.Setenv(modeEnvVarname, devMode)
	defer func() {
		_ = os.Unsetenv(modeEnvVarname)
	}()
	if err := SetupStores(); err != nil {
		t.Error(err)
	}
	_, ok := GlobalTrafficJamStore.(*store.InMemoryTrafficJamStore)
	if !ok {
		t.Errorf("expected an in memory store pointer")
	}
}

func TestSetupStores_Prod(t *testing.T) {
	_ = os.Setenv(modeEnvVarname, prodMode)
	defer func() {
		_ = os.Unsetenv(modeEnvVarname)
	}()
	if err := SetupStores(); err == nil {
		t.Error("expected an error")
	}
}
