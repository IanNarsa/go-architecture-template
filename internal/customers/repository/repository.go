package repository

import "go-arch/internal/customers/model"

type CustomerRepository interface {
	GetAllCustomer() ([]model.Customer, error)
}
