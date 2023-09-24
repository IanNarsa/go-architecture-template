package handler

import (
	"encoding/json"
	"go-arch/internal/products/model"
	"go-arch/internal/products/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func (h *ProductHandler) GetAllProduct(c echo.Context) error {
	products, err := h.productUsecase.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetSelectedProduct(c echo.Context) error {

	productCode := c.Param("productCode")

	proeductSelected, err := h.productUsecase.GetSelectedProduct(model.SelectedProduct{ProductCode: productCode})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, proeductSelected)
}

func (h *ProductHandler) OrderProduct(c echo.Context) error {
	var payload model.OrderProducts

	// Decode the JSON payload from the request body
	if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		// Handle the error, e.g., return an error response
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
	}

	// Call the OrderProduct method with the decoded payload
	if err := h.productUsecase.OrderProduct(payload); err != nil {
		// Handle the error, e.g., return an error response
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	// Handle a successful order, e.g., return a success response
	return c.JSON(http.StatusOK, map[string]string{"message": "Order placed successfully"})
}
