package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customers struct {
	Name    string `json:"full_name" xml:"Name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customers{
		{Name: "Raj", City: "Sasaram", ZipCode: "821113"},
		{Name: "MRaj", City: "Hyderabad", ZipCode: "500075"},
	}

	
	if r.Header.Get("Content-type") == "application/xml"{
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}else{
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
	
}