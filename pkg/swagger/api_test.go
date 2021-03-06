package swagger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/peterjochum/traffic-jam-api/internal/app"
	"github.com/peterjochum/traffic-jam-api/pkg/models"
	"github.com/peterjochum/traffic-jam-api/pkg/store"

	"github.com/gorilla/mux"
)

const trafficJamAPIRoot = "/api/v1/trafficjam/"

// APITestCase encapsulates test input/output
type APITestCase struct {
	name           string
	id             string
	expectedStatus int
	body           string
}

type deleteAPITestCase struct {
	APITestCase
	expectedJamCount int64
}

func getTestTrafficJamBody(id int) string {
	switch id {
	case 1:
		// this jam is inserted by the SeedTrafficJamStore function
		return "{\"id\":1,\"longitude\":1.12,\"latitude\":2.13,\"durationInSeconds\":15}\n"
	case 4:
		// new traffic jam
		return "{\"id\":4,\"longitude\":1.15,\"latitude\":2.15,\"durationInSeconds\":60}\n"
	default:
		return ""
	}

}

func TestGetAllTrafficJams(t *testing.T) {
	app.GlobalTrafficJamStore = store.NewInMemoryTrafficJamStore()
	store.SeedTrafficJamStore(app.GlobalTrafficJamStore)

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
	app.GlobalTrafficJamStore = store.NewInMemoryTrafficJamStore()
	store.SeedTrafficJamStore(app.GlobalTrafficJamStore)
	handler := http.HandlerFunc(DeleteTrafficJam)

	testcases := []deleteAPITestCase{
		{APITestCase: APITestCase{
			name:           "nonexisting",
			id:             "4",
			expectedStatus: http.StatusNotFound,
			body:           "",
		}, expectedJamCount: 3},
		{
			APITestCase: APITestCase{
				name:           "illegal id",
				id:             "abc",
				expectedStatus: http.StatusBadRequest,
				body:           "",
			},
			expectedJamCount: 3,
		},
		{
			APITestCase: APITestCase{
				name:           "sucessful delete",
				id:             "1",
				expectedStatus: http.StatusOK,
				body:           "",
			},
			expectedJamCount: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", trafficJamAPIRoot+tc.id, nil)
			rr := httptest.NewRecorder()

			// Fake the parameter parsing
			vars := map[string]string{
				"id": tc.id,
			}
			req = mux.SetURLVars(req, vars)
			handler.ServeHTTP(rr, req)

			if tc.expectedStatus != rr.Code {
				t.Errorf("Expected status %d but got %d",
					tc.expectedStatus, rr.Code)
			}

			if tc.expectedJamCount != app.GlobalTrafficJamStore.Total() {
				t.Errorf("Expected %d jams left, but there are %d",
					tc.expectedJamCount, app.GlobalTrafficJamStore.Total())
			}
		})
	}
}

func TestAddTrafficJam(t *testing.T) {
	msgMalformedBody := "could not parse traffic jam from request body"
	cases := []struct {
		name        string
		body        string
		eResultCode int
		eMessage    string
		totalJams   int64
	}{
		{"empty body", "", http.StatusBadRequest, msgMalformedBody, 3},
		{"malformed body", "malformed", http.StatusBadRequest, msgMalformedBody, 3},
		{"object exists", getTestTrafficJamBody(1), http.StatusUnprocessableEntity, "traffic jam 1 already exists", 3},
		{"successful add", getTestTrafficJamBody(4), http.StatusOK, "success", 4},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			app.GlobalTrafficJamStore = store.NewInMemoryTrafficJamStore()
			store.SeedTrafficJamStore(app.GlobalTrafficJamStore)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(AddTrafficJam)
			bodyReader := strings.NewReader(test.body)
			req := httptest.NewRequest("POST", trafficJamAPIRoot, bodyReader)
			handler.ServeHTTP(rr, req)

			if rr.Code != test.eResultCode {
				t.Errorf("Expected result code %d but got %d",
					test.eResultCode, rr.Code)
			}

			if rr.Body.String() != test.eMessage {
				t.Errorf("Expected response \"%s\" but got \"%s\"",
					test.eMessage, rr.Body.String())
			}

			if test.totalJams != app.GlobalTrafficJamStore.Total() {
				t.Errorf("expected %d jams but got %d", test.totalJams, app.GlobalTrafficJamStore.Total())
			}

		})
	}

}

func TestGetTrafficJam(t *testing.T) {
	app.GlobalTrafficJamStore = store.NewInMemoryTrafficJamStore()
	store.SeedTrafficJamStore(app.GlobalTrafficJamStore)
	handler := http.HandlerFunc(GetTrafficJam)

	testcases := []APITestCase{
		{name: "nonexisting", id: "99", expectedStatus: http.StatusNotFound, body: "object not found"},
		{name: "wrongid", id: "abcd", expectedStatus: http.StatusBadRequest, body: "unable to parse id"},
		{name: "existing", id: "1", expectedStatus: http.StatusOK, body: getTestTrafficJamBody(1)},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			target := fmt.Sprintf("%s%s", trafficJamAPIRoot, tc.id)
			req := httptest.NewRequest("GET", target, nil)

			// Fake the parameter parsing
			vars := map[string]string{
				"id": tc.id,
			}

			req = mux.SetURLVars(req, vars)
			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("got status %d, expected %d",
					rr.Code, tc.expectedStatus)
			}

			if strings.Compare(rr.Body.String(), tc.body) != 0 {
				t.Errorf("expected body:\n%s\ngot:\n%s",
					tc.body, rr.Body.String())
			}
		})
	}

}

func TestPutTrafficJam(t *testing.T) {
	app.GlobalTrafficJamStore = store.NewInMemoryTrafficJamStore()
	store.SeedTrafficJamStore(app.GlobalTrafficJamStore)

	handler := http.HandlerFunc(PutTrafficJam)
	updatedJam := getTestTrafficJamBody(1)

	testCases := []APITestCase{
		{"bad id", "abc", http.StatusBadRequest, ""},
		{"nonexisting jam", "4", http.StatusNotFound, getTestTrafficJamBody(4)},
		{"invalid body", "1", http.StatusBadRequest, "{:sdfsdf:SDFsdf}"},
		{"success", "1", http.StatusOK, updatedJam},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			target := fmt.Sprintf("%s%s", trafficJamAPIRoot, tc.id)
			bodyReader := strings.NewReader(tc.body)
			req := httptest.NewRequest("PUT", target, bodyReader)
			// Fake the parameter parsing
			vars := map[string]string{
				"id": tc.id,
			}
			req = mux.SetURLVars(req, vars)

			handler.ServeHTTP(rr, req)

			if tc.expectedStatus != rr.Code {
				t.Errorf("expected code %d, got %d",
					tc.expectedStatus, rr.Code)
			}
		})
	}
}
