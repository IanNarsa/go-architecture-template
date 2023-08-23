package usecase

import "go-arch/pkg/customers/model"

type CustomerUsecase interface {
	GetAllCustomer() (*[]model.Customer, error)
}
