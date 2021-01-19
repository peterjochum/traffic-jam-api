package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex_WrongStaticPath(t *testing.T) {
	handler := http.HandlerFunc(Index)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	// Handle panic
	defer func() {
		_ = recover()
	}()

	// Assume a panic happens here
	handler.ServeHTTP(rr, req)
	t.Error("did not panic")
}

func TestIndex(t *testing.T) {
	handler := http.HandlerFunc(Index)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	StaticFilesPath = "../../static"
	handler.ServeHTTP(rr, req)
	needles := []string{"Github", "Swagger", "Traffic Jam API", "/api/v1/trafficjam"}
	for _, needle := range needles {
		if !strings.Contains(rr.Body.String(), needle) {
			t.Errorf("expected \"%s\" in response", needle)
		}
	}
}
