package usecase

import "go-arch/internal/customers/model"

type CustomerUsecase interface {
	GetAllCustomer() (*[]model.Customer, error)
}
