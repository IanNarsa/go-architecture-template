package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository_GetAllCustomer(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewCustomerRepository(db)

	rows := sqlmock.NewRows([]string{"customerNumber", "customerName", "phone"}).
		AddRow(1, "John Doe", "123-456-7890").
		AddRow(2, "Jane Smith", "987-654-3210")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	customers, err := repo.GetAllCustomer()
	assert.NoError(t, err)
	assert.Len(t, customers, 2)
}

func TestCustomerRepository_GetAllCustomer_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewCustomerRepository(db)

	mock.ExpectQuery("SELECT").WillReturnError(errors.New("database error"))

	customers, err := repo.GetAllCustomer()
	assert.Error(t, err)
	assert.Nil(t, customers)
}
