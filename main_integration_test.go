/*
Created on Mon May 16 17:08:23 2022

@author: amolzinjade@gmail.com
*/
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
	q.Add("EmailAddress", "amol@gmail.com")
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
	q.Add("LastName", "Zinjade")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(searchByEmail)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
