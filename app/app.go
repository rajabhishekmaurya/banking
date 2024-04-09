package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rajabhishekmaurya/banking/domain"
	"github.com/rajabhishekmaurya/banking/service"
)

// func sanitycheck() {
// 	if os.Getenv("SERVER_ADDRESS") == "" ||
// 		os.Getenv("SERVER_PORT") == "" {
// 		log.Fatal("Environment variable not define....")
// 	}
// }

func Start() {

	//sanitycheck()

	router := mux.NewRouter()
	//wiring............
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomeRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomeRepositoryDb())}

	//define routes

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS") // set env. var in terminal like this: SERVER_ADDRESS=localhost SERVER_PORT=8000 go run main.go
	port := os.Getenv("SERVER_PORT")

	//starting the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
