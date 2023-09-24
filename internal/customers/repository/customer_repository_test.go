package repository

import (
	"testing"

	"go-arch/internal/customers/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllCustomer(t *testing.T) {
	// Create a new SQL mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Create a new CustomerImpl instance with the mock database
	repo := &CustomerImpl{db: db}

	// Define the expected rows and columns
	rows := sqlmock.NewRows([]string{"customerNumber", "customerName", "phone"}).
		AddRow(1, "Customer 1", "123-456-7890").
		AddRow(2, "Customer 2", "987-654-3210")

	// Set up the mock expectations
	mock.ExpectQuery("select customerNumber, customerName, phone from customers").WillReturnRows(rows)

	// Call the method you want to test
	customers, err := repo.GetAllCustomer()
	if err != nil {
		t.Errorf("Error calling GetAllCustomer: %v", err)
	}

	// Define the expected result
	expectedCustomers := []model.Customer{
		{CustomerNumber: 1, CustomerName: "Customer 1", Phone: "123-456-7890"},
		{CustomerNumber: 2, CustomerName: "Customer 2", Phone: "987-654-3210"},
	}

	// Compare the actual result to the expected result
	if len(customers) != len(expectedCustomers) {
		t.Errorf("Expected %d customers, but got %d", len(expectedCustomers), len(customers))
	}

	for i, actual := range customers {
		expected := expectedCustomers[i]
		if actual != expected {
			t.Errorf("Mismatch in customer data at index %d\nExpected: %+v\nActual:   %+v", i, expected, actual)
		}
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
