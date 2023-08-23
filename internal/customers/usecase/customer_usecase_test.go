package usecase

import (
	"errors"
	"testing"

	"go-arch/internal/customers/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCustomerRepository struct {
	mock.Mock
}

func (m *mockCustomerRepository) GetAllCustomer() ([]model.Customer, error) {
	args := m.Called()
	return args.Get(0).([]model.Customer), args.Error(1)
}

func TestCustomerUsecase_GetAllCustomer(t *testing.T) {
	repo := &mockCustomerRepository{}
	usecase := NewCustomerUsecase(repo)

	expectedCustomers := []model.Customer{
		{CustomerNumber: 1, CustomerName: "Customer A", Phone: "123-456"},
		{CustomerNumber: 2, CustomerName: "Customer B", Phone: "789-012"},
	}

	repo.On("GetAllCustomer").Return(expectedCustomers, nil)

	customers, err := usecase.GetAllCustomer()
	assert.NoError(t, err)
	assert.Len(t, customers, 2)
	assert.Equal(t, expectedCustomers, customers)

	repo.AssertExpectations(t)
}

func TestCustomerUsecase_GetAllCustomer_Error(t *testing.T) {
	repo := new(mockCustomerRepository)
	usecase := NewCustomerUsecase(repo)

	expectedError := errors.New("repository error")
	repo.On("GetAllCustomer").Return(nil, expectedError)

	customers, err := usecase.GetAllCustomer()
	assert.Error(t, err)
	assert.Nil(t, customers)
	assert.Equal(t, expectedError, err)

	repo.AssertExpectations(t)
}
