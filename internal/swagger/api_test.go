package swagger

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/peterjochum/traffic-jam-api/internal/app"
	"github.com/peterjochum/traffic-jam-api/internal/models"
	"github.com/peterjochum/traffic-jam-api/internal/store"
)

func TestGetAllTrafficJams(t *testing.T) {
	app.TrafficJamStore = store.NewInMemoryTrafficJamStore(true)

	req := httptest.NewRequest("GET", "/api/v1/trafficjam", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllTrafficJams)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d", http.StatusOK)
	}

	eContentType := "application/json; charset=UTF-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != eContentType {
		t.Errorf("Got content type %v expected %v", contentType, eContentType)
	}

	var trafficJamResponse []models.TrafficJam
	err := json.NewDecoder(rr.Body).Decode(&trafficJamResponse)
	if err != nil {
		t.Error(err)
	}

	eCountJams := 3
	actualJams := len(trafficJamResponse)
	if actualJams != eCountJams {
		t.Errorf("Expected %d traffic jams, but got %d",
			eCountJams, actualJams)
	}
}

func TestDeleteTrafficJam(t *testing.T) {
	app.TrafficJamStore = store.NewInMemoryTrafficJamStore(true)

	req := httptest.NewRequest("DELETE", "/api/v1/trafficjam/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTrafficJam)

	// Fake the parameter parsing
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d", http.StatusOK)
	}

	remainingJams := app.TrafficJamStore.ListTrafficJams()
	eJams := 2
	aJams := len(remainingJams)
	if aJams != eJams {
		t.Errorf("Expected %d jams left, but there are %d", eJams, aJams)
	}
}

func TestAddTrafficJam(t *testing.T) {

}

func TestGetTrafficJam(t *testing.T) {

}

func TestPutTrafficJam(t *testing.T) {

}

func TestIndex(t *testing.T) {

}
