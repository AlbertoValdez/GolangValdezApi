package main

import (
	"database/sql"
	"net/http"

	"github.com/GolangValdezApi/databasecon"
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

	var productRepository = product.Nr(dbco)
	var productService product.Service
	productService = product.Ns(productRepository)
	r := chi.NewRouter()
	r.Mount("/products", product.MakeHTTPHandler(productService))
	http.ListenAndServe(":3000", r)
}
