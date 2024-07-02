package repository

import (
	"database/sql"
	"fmt"
	"go-arch/internal/products/model"
	"go-arch/pkg/logger"
	"time"
)

// ProductImpl for implementation on function
type ProductImpl struct {
	db *sql.DB
}

// create func get all products
func (r *ProductImpl) GetAllProducts() (*[]model.Products, error) {
	var products []model.Products

	rows, err := r.db.Query("SELECT productCode, productName, productVendor, quantityInStock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := model.Products{}

		if err := rows.Scan(&product.ProductCode, &product.ProductName, &product.ProductVendor, &product.Stock); err != nil {
			logger.Error(err)
			return nil, err
		}
		products = append(products, product)

	}

	if err := rows.Err(); err != nil {
		logger.Error(err)
		return nil, err
	}

	return &products, nil
}

func (r *ProductImpl) GetSelectedProduct(code model.SelectedProduct) (*[]model.Products, error) {
	var products []model.Products

	rows, err := r.db.Query("select productCode, productName, productVendor, quantityInStock FROM products where productCode = '" + code.ProductCode + "';")
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := model.Products{}

		if err := rows.Scan(
			&product.ProductCode,
			&product.ProductName,
			&product.ProductVendor,
			&product.Stock,
		); err != nil {
			logger.Error(err)
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		logger.Error(err)
		return nil, err
	}

	return &products, nil
}

func (r *ProductImpl) OrderProducts(data model.OrderProducts) error {

	dt := &model.OrderProducts{
		ProductCode:    data.ProductCode,
		Quantity:       data.Quantity,
		CustomerNumber: data.CustomerNumber,
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading time zone:", err)
	}

	// Get the current time in the "Asia/Jakarta" time zone
	currentTime := time.Now().In(loc)

	tx, err := r.db.Begin()
	if err != nil {

		logger.Error(err)
	}

	// Insert into the "orders" table
	order, err := tx.Exec(`INSERT INTO orders (
		orderDate, requiredDate,
		status, comments, 
		customerNumber) VALUES (?, ?,?, ?, ?)`,
		currentTime.Format("2006-01-02"), currentTime.Format("2006-01-02"), "waiting", nil, dt.CustomerNumber)
	if err != nil {
		tx.Rollback()
		logger.Error(err)
	}

	id, _ := order.LastInsertId() // get last insert id from database and assign it to variable 'id' for further use later

	var orderDetail model.OrderDetail

	err = r.db.QueryRow(`SELECT MSRP from products where productCode = ?`, dt.ProductCode).Scan(&orderDetail.Price)
	if err != nil {
		tx.Rollback()

		logger.Error(err)
	}

	// Insert into the "orderdetails" table
	_, err = tx.Exec(`INSERT INTO orderdetails (
		orderNumber, productCode, 
		quantityOrdered, priceEach, 
		orderLineNumber) VALUES (?, ?, ?, ?, ?)`,
		id, dt.ProductCode, dt.Quantity, orderDetail.Price, 0)

	if err != nil {
		tx.Rollback()

		logger.Error(err)

	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()

		logger.Error(err)
	}

	logger.Info("Success create order")
	return err
}
