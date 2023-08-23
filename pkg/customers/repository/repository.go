package repository

import "go-arch/pkg/customers/model"

type CustomerRepository interface {
	GetAllCustomer() (*[]model.Customer, error)
}
