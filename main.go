package main

import (
	"fmt"
	"go-arch/config"
	"go-arch/pkg/customers/handler"
	"go-arch/pkg/customers/repository"
	"go-arch/pkg/customers/usecase"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDatabase("localhost:3306", "company_a", "root", "root")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()

	// Create customer repository
	customerRepo := repository.NewCustomerRepository(db)

	// Create customer usecase
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)

	// Register customer handler
	handler.RegisterCustomerHandler(e, customerUsecase)

	// Start the server
	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))

}
