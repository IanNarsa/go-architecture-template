package handler

import (
	"go-arch/internal/customers/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterCustomerHandler(e *echo.Echo, customerUsecase usecase.CustomerUsecase) {
	handler := NewCustomerHandler(customerUsecase)

	e.GET("/customers", handler.GetAllCustomer)
}
