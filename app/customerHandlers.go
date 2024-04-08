package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajabhishekmaurya/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customers{
	// 	{Name: "Raj", City: "Sasaram", ZipCode: "821113"},
	// 	{Name: "MRaj", City: "Hyderabad", ZipCode: "500075"},
	// }

	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)

	if err !=nil {
		writeResponse(w, err.Code, err.Message)
	}else{
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}else{
		writeResponse(w, http.StatusOK, customer)
	}
	
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err !=nil {
		panic(err)
	}

}