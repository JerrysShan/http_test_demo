package handlers_test

import (
	"encoding/json"
	"http_test_demo/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	handlers.Routes()
}

func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request", ballotX, err)
		}
		t.Log("\tShould be able to create a request", checkMark)
		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)
		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould Receive \"200\"")
		u := struct {
			Name  string
			Email string
		}{}
		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould Decode the response", ballotX)
		}
		t.Log("\tShould Decode the response", checkMark)
	}
}
