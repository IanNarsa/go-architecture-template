package handler

import (
	"go-arch/internal/products/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterProductHandler(e *echo.Echo, productUsecase usecase.ProductUsecase) {
	handler := NewProductHandler(productUsecase)

	e.GET("/products", handler.GetAllProduct)
	e.GET("/product/:productCode", handler.GetSelectedProduct)
	e.POST("/order-product", handler.OrderProduct)
}
