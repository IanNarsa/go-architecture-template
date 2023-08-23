package usecase

import (
	"go-arch/pkg/customers/model"
	"go-arch/pkg/customers/repository"
)

type CustomerUsecaseImpl struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerUsecase(customerRepository repository.CustomerRepository) CustomerUsecase {
	return &CustomerUsecaseImpl{customerRepository: customerRepository}
}

func (u *CustomerUsecaseImpl) GetAllCustomer() (*[]model.Customer, error) {
	return u.customerRepository.GetAllCustomer()
}
