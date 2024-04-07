package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/rajabhishekmaurya/banking/service"
)

type Customers struct {
	Name    string `json:"full_name" xml:"Name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customers{
	// 	{Name: "Raj", City: "Sasaram", ZipCode: "821113"},
	// 	{Name: "MRaj", City: "Hyderabad", ZipCode: "500075"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	
	if r.Header.Get("Content-type") == "application/xml"{
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}else{
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
	
}