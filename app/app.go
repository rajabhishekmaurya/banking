package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rajabhishekmaurya/banking/domain"
	"github.com/rajabhishekmaurya/banking/logger"
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
	dbClient := getDbClient()

	customerRepositoryDb :=  domain.NewCustomeRepositoryDb(dbClient)
	accountRepositoryDb :=  domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	//define routes

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)


	address := os.Getenv("SERVER_ADDRESS") // set env. var in terminal like this: SERVER_ADDRESS=localhost SERVER_PORT=8000 go run main.go
	port := os.Getenv("SERVER_PORT")

	//starting the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {

	username := "root"
	password := "my-Password"
	host := "179.100.0.2"
	port := "3306"
	databaseName := "banking"

	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWD")
	// host := os.Getenv("DB_ADDR")
	// port := os.Getenv("DB_PORT")
	// databaseName := os.Getenv("DB_NAME")

	// Create a connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, databaseName)

	client, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	logger.Info("Connected to the MySQL database!")

	return client
}
