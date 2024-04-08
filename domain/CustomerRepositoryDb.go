package domain

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rajabhishekmaurya/banking/errs"
	"github.com/rajabhishekmaurya/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}
func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)

		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	

	if err != nil {
		logger.Error("Error while scanning customers " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError){
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer

	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}else{
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}

	}
	return &c, nil
}

func NewCustomeRepositoryDb() CustomerRepositoryDb {

	username := "root"
	password := "my-Password"
	host := "179.100.0.2"
	port := "3306"
	databaseName := "banking"

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

	return CustomerRepositoryDb{client}

}