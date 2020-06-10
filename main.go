package main

import (
	"database/sql"
	"net/http"

	"github.com/GolangValdezApi/costumers"
	"github.com/GolangValdezApi/databasecon"
	"github.com/GolangValdezApi/employee"
	"github.com/GolangValdezApi/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

//Product estructura
type Product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productcode"`
	Description string `json:"description"`
}

var dbco *sql.DB

func main() {

	dbco = databasecon.InitDB()
	defer dbco.Close()

	var (
		productRepository  = product.Nr(dbco)
		employeeRepository = employee.NewRepository(dbco)
		costumerRepository = costumers.NewRepository(dbco)
	)

	var (
		productService  product.Service
		employeeService employee.Service
		costumerService costumers.Service
	)

	productService = product.Ns(productRepository)
	employeeService = employee.NewService(employeeRepository)
	costumerService = costumers.NewService(costumerRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHTTPHandler(productService))
	r.Mount("/empleados", employee.MakeHTTPHandler(employeeService))
	r.Mount("/clientes", costumers.MakeHTTPHandler(costumerService))

	http.ListenAndServe(":3000", r)
}
