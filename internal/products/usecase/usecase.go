package usecase

import (
	"go-arch/internal/products/model"
	"go-arch/internal/products/repository"
)

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &ProductUsecaseImpl{productRepository: productRepository}
}

type ProductUsecase interface {
	GetAllProducts() (*[]model.Products, error)
	GetSelectedProduct(param model.SelectedProduct) (*[]model.Products, error)
	OrderProduct(data model.OrderProducts) error
}
