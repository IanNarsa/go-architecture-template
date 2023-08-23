package handler

import (
	"net/http"

	"go-arch/internal/customers/usecase"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerHandler(customerUsecase usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{customerUsecase: customerUsecase}
}

func (h *CustomerHandler) GetAllCustomer(c echo.Context) error {
	customers, err := h.customerUsecase.GetAllCustomer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, customers)
}

func RegisterCustomerHandler(e *echo.Echo, customerUsecase usecase.CustomerUsecase) {
	handler := NewCustomerHandler(customerUsecase)

	e.GET("/customers", handler.GetAllCustomer)
}
