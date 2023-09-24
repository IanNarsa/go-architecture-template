package repository

import (
	"database/sql"
	"go-arch/internal/products/model"
)

type ProductRepository interface {
	GetAllProducts() (*[]model.Products, error)
	GetSelectedProduct(param model.SelectedProduct) (*[]model.Products, error)
	OrderProducts(data model.OrderProducts) error
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductImpl{db: db}
}
