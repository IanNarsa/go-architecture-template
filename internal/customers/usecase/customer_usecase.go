package usecase

import (
	"go-arch/internal/customers/model"
	"go-arch/internal/customers/repository"
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
