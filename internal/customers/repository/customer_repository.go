package repository

import (
	"database/sql"
	"go-arch/internal/customers/model"
)

type CustomerImpl struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &CustomerImpl{db: db}
}

func (r *CustomerImpl) GetAllCustomer() (*[]model.Customer, error) {

	query := "select customerNumber, customerName, phone from customers"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.CustomerNumber, &customer.CustomerName, &customer.Phone); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &customers, nil
}
