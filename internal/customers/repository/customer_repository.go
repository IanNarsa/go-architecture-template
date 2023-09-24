package repository

import (
	"database/sql"
	"go-arch/internal/customers/model"
	"log"
)

type CustomerImpl struct {
	db *sql.DB
}

func (r *CustomerImpl) GetAllCustomer() ([]model.Customer, error) {

	query := "select customerNumber, customerName, phone from customers"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("[Error] ", err)
		return nil, err
	}
	defer rows.Close()

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.CustomerNumber, &customer.CustomerName, &customer.Phone); err != nil {
			log.Println("[Error] ", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		log.Println("[Error] ", err)
		return nil, err
	}

	return customers, nil
}
