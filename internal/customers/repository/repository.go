package repository

import (
	"database/sql"
	"go-arch/internal/customers/model"
)

type CustomerRepository interface {
	GetAllCustomer() ([]model.Customer, error)
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &CustomerImpl{db: db}
}
