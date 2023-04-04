package antena

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	PingHandler(rr, request)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected %v got %v", http.StatusOK, rr.Result().StatusCode)
	}
	var res Response
	err := json.NewDecoder(rr.Body).Decode(&res)
	if err != nil {
		t.Fatal(err)
	}
	if res.Text != "pong" {
		t.Errorf("Expected %v got %v", "pong", res.Text)
	}
}
