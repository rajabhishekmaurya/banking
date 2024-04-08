package service

import (
	"github.com/rajabhishekmaurya/banking/domain"
	"github.com/rajabhishekmaurya/banking/dto"
	"github.com/rajabhishekmaurya/banking/errs"
)


type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer,*errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError){
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError){
	c, err := s.repo.ById(id)
	if err !=nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}