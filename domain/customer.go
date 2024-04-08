package domain

import (
	"github.com/rajabhishekmaurya/banking/errs"
)


type Customer struct {
	Id 			string `db:"customer_id"`
	Name 		string
	City 		string
	Zipcode 	string
	DateofBirth string	`db:"date_of_birth"`
	Status 		string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
