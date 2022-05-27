/*
Created on Mon May 16 17:08:23 2022

@author: amolzinjade@gmail.com
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

type contact struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	EmailAddress string `json:"EmailAddress"`
}

type allContacts []contact

var contacts = allContacts{
	{
		
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createContact(w http.ResponseWriter, r *http.Request) {
	var newContact contact
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Email Address only in order to update")
	}
	
	json.Unmarshal(reqBody, &newContact)
	contacts = append(contacts, newContact)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newContact)
}

func searchByEmail(w http.ResponseWriter, r *http.Request) {
	contactEmailAddress := mux.Vars(r)["EmailAddress"]

	for _, singleContact := range contacts {
		if singleContact.EmailAddress == contactEmailAddress {
			json.NewEncoder(w).Encode(singleContact)
		}
	}
}
func searchByLastName(w http.ResponseWriter, r *http.Request) {
	contactLastName := mux.Vars(r)["LastName"]

	for _, singleContact := range contacts {
		if strings.Contains(singleContact.LastName, contactLastName) {
			json.NewEncoder(w).Encode(singleContact)
		}
	}
}
func getAllContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contacts)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	contactEmailAddress := mux.Vars(r)["EmailAddress"]
	var updatedContact contact

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with Email Address only in order to update")
	}
	json.Unmarshal(reqBody, &updatedContact)

	for i, singleContact := range contacts {
		if singleContact.EmailAddress == contactEmailAddress {
			singleContact.FirstName = updatedContact.FirstName
			singleContact.LastName = updatedContact.LastName
			contacts = append(contacts[:i], singleContact)
			json.NewEncoder(w).Encode(singleContact)
		}
	}
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	contactEmailAddress := mux.Vars(r)["EmailAddress"]

	for i, singleContact := range contacts {
		if singleContact.EmailAddress == contactEmailAddress {
			contacts = append(contacts[:i], contacts[i+1:]...)
			fmt.Fprintf(w, "The contact with Email Address %v has been deleted successfully", contactEmailAddress)
		}
	}
}

func main() {
	fmt.Println("Starting server on port:8080")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/contact", createContact).Methods("POST")
	router.HandleFunc("/contacts", getAllContacts).Methods("GET")
	router.HandleFunc("/contacts/{EmailAddress}", searchByEmail).Methods("GET")
	router.HandleFunc("/contacts/{LastName}", searchByLastName).Methods("GET")
	router.HandleFunc("/contacts/{EmailAddress}", updateContact).Methods("PATCH")
	router.HandleFunc("/contacts/{EmailAddress}", deleteContact).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
