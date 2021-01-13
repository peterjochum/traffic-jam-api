package swagger

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/peterjochum/traffic-jam-api/pkg/app"
	"github.com/peterjochum/traffic-jam-api/pkg/models"
	"github.com/peterjochum/traffic-jam-api/pkg/store"

	"github.com/gorilla/mux"
)

const trafficJamAPIRoot = "/api/v1/trafficjam/"

func TestGetAllTrafficJams(t *testing.T) {
	app.TrafficJamStore = store.NewInMemoryTrafficJamStore(true)

	req := httptest.NewRequest("GET", trafficJamAPIRoot, nil)
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

	req := httptest.NewRequest("DELETE", trafficJamAPIRoot+"1", nil)
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
	cases := []struct {
		name        string
		body        io.Reader
		eResultCode int
		eMessage    string
	}{
		{"empty body", nil, 400, "could not parse trafficjam from request body"},
		// TODO: Test malformed body
		// TODO: Test object exists
		// TODO: Test success
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			app.TrafficJamStore = store.NewInMemoryTrafficJamStore(true)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(AddTrafficJam)
			req := httptest.NewRequest("POST", trafficJamAPIRoot, nil)
			handler.ServeHTTP(rr, req)

			if rr.Code != test.eResultCode {
				t.Errorf("Expected result code %d but got %d",
					test.eResultCode, rr.Code)
			}

			if rr.Body.String() != test.eMessage {
				t.Errorf("Expected response \"%s\" but got \"%s\"",
					test.eMessage, rr.Body.String())
			}
		})
	}

}

func TestGetTrafficJam(t *testing.T) {

}

func TestPutTrafficJam(t *testing.T) {

}

func TestIndex(t *testing.T) {

}
