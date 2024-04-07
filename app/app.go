package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajabhishekmaurya/banking/domain"
	"github.com/rajabhishekmaurya/banking/service"
)


func Start(){
	// mux := http.NewServeMux()

	router := mux.NewRouter()
	//wiring............
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomeRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomeRepositoryDb())}


	//define routes
	
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	
	//starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
