package swagger

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("testing"))
}

func TestLogger(t *testing.T) {
	handler := http.HandlerFunc(testHandler)
	l := Logger(handler, "Test")
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	l.ServeHTTP(rr, req)

	if rr.Body.String() != "testing" {
		t.Errorf("logger should leave original response intact")
	}

	// NTH: check the logged data
}

func TestCorsHeaderSetter(t *testing.T) {
	handler := http.HandlerFunc(testHandler)
	testOrigin := "*"
	handlerWithCors := CorsHeaderSetter(handler, testOrigin)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	handlerWithCors.ServeHTTP(rr, req)

	if rr.Header().Get("Access-Control-Allow-Origin") != testOrigin {
		t.Errorf("access control header missing on request")
	}
}
