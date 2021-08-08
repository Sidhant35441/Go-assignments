package service

import (
	"gorestapi/entity"
	"gorestapi/repository"
)

type CustomerService interface {
	Save(cust entity.Customer) error
	FindAll() []entity.Customer
	Update(cust entity.Customer) error
	Delete(cust entity.Customer) error
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

//creating new customer
func (service *customerService) Save(cust entity.Customer) error {
	service.customerRepository.Save(cust)
	return nil
}

//find all customers
func (service *customerService) FindAll() []entity.Customer {
	return service.customerRepository.FindAll()
}

//Update customer
func (service *customerService) Update(cust entity.Customer) error {
	service.customerRepository.Update(cust)
	return nil
}

//Delete customer
func (service *customerService) Delete(cust entity.Customer) error {
	service.customerRepository.Delete(cust)
	return nil
}
