package main
import (
	"testing"
	"net/http"
	"net/http/httptest"
)
func TestGetContactByEmailId(t *testing.T) {

	req, err := http.NewRequest("GET", "/contacts", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("EmailAddress", "hammad@gmail.com")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(searchByEmail)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
func TestGetContactByLastName(t *testing.T) {

	req, err := http.NewRequest("GET", "/contacts", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("LastName", "Mehdi")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(searchByEmail)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}