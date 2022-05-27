/*
Created on Mon May 16 17:08:23 2022

@author: amolzinjade@gmail.com
*/
package main
import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
)

func TestCreateContact(t *testing.T) {

	var jsonStr = []byte(`{"FirstName":"Amol","LastName":"Zinjade","EmailAddress":"amol@gmail.com"}`)

	req, err := http.NewRequest("POST", "/contact", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createContact)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expected := rr.Body.String()
	
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestAllContacts(t *testing.T) {
	req, err := http.NewRequest("GET", "/contacts", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllContacts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdateContact(t *testing.T) {

	var jsonStr = []byte(`{"FirstName":"Amol","LastName":"Zinjade","EmailAddress":"amol@gmail.com"}`)

	req, err := http.NewRequest("PATCH", "/contacts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("EmailAddress", "amol@gmail.com")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateContact)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := rr.Body.String()
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestDeleteEntry(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/contacts", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("EmailAddress", "amol@gmail.com")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteContact)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := rr.Body.String()
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
