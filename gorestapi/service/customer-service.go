package service

import (
	"gorestapi/entity"
	"gorestapi/repository"
)

type CustomerService interface {
	Save(cust entity.Customer) entity.Customer
	FindAll() []entity.Customer
}
type customerService struct {
	//customers []entity.Customer
	customerRepository repository.CustomerRepository
}

func New(repo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: repo,
	}
}
func (service *customerService) Save(cust entity.Customer) entity.Customer {
	service.customerRepository.Save(cust)
	return cust
}
func (service *customerService) FindAll() []entity.Customer {
	return service.customerRepository.FindAll() // actual &service.customers
}
