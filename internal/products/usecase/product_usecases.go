package usecase

import (
	"go-arch/internal/products/model"
	"go-arch/internal/products/repository"
)

// struct get repository
type ProductUsecaseImpl struct {
	productRepository repository.ProductRepository
}

func (u *ProductUsecaseImpl) GetAllProducts() (*[]model.Products, error) {
	return u.productRepository.GetAllProducts()
}

func (u *ProductUsecaseImpl) GetSelectedProduct(param model.SelectedProduct) (*[]model.Products, error) {
	return u.productRepository.GetSelectedProduct(param)
}

func (u *ProductUsecaseImpl) OrderProduct(data model.OrderProducts) error {
	return u.productRepository.OrderProducts(data)
}
