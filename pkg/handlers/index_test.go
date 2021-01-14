package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {

	handler := http.HandlerFunc(Index)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	handler.ServeHTTP(rr, req)
	if rr.Body.String() != apiHelp {
		t.Errorf("unexpected body:\n%s\nexpected:\n%s",
			rr.Body.String(), apiHelp)
	}

}
