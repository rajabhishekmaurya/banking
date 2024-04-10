package domain

import (
	"github.com/rajabhishekmaurya/banking/dto"
	"github.com/rajabhishekmaurya/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId:a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
