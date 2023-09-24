package model

type Products struct {
	ProductCode   string
	ProductName   string
	ProductVendor string
	Stock         int
}

type SelectedProduct struct {
	ProductCode string
}

type OrderProducts struct {
	OrderNumber    string
	ProductCode    string
	Quantity       int
	CustomerNumber int64
}

type OrderDetail struct {
	ProductID int
	Quantity  int
	Price     float64
}
