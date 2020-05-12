package product

import (
	"database/sql"

	"github.com/GolangValdezApi/helper"
)

//Repository interface
type Repository interface {
	GetProductByID(productID int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProducts(params *getAddProductRequest) (int64, error)
	UpdateProducts(params *updateProductRequest) (int64, error)
	DeleteProducts(params *deleteProductRequest) (int64, error)
	getBestSellers() ([]*ProductTop, error)
	getTotalVentas() (float64, error)
}

type repository struct {
	db *sql.DB
}

//Nr bd
func Nr(dbco *sql.DB) Repository {
	return &repository{db: dbco}
}

//Func adicional
func (repo *repository) GetProductByID(productID int) (*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),standard_cost,list_price,category
				FROM products
				WHERE id=?`
	row := repo.db.QueryRow(sql, productID)
	product := &Product{}
	err := row.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)
	helper.Catch(err)

	return product, err
}

func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),
					standard_cost,list_price,category
					FROM products
				LIMIT ? OFFSET ?`

	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	if err != nil {
		panic(err)
	}

	var product []*Product
	for results.Next() {
		products := &Product{}
		err = results.Scan(&products.ID, &products.ProductCode, &products.ProductName, &products.Description, &products.StandardCost, &products.ListPrice, &products.Category)
		if err != nil {
			panic(err)
		}

		product = append(product, products)
	}

	return product, nil
}

func (repo *repository) GetTotalProducts() (int, error) {
	const sql = "SELECT COUNT(*) FROM products"
	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}

func (repo *repository) InsertProducts(params *getAddProductRequest) (int64, error) {
	const sql = `INSERT INTO products (product_code, product_name,category,description,list_price,standard_cost)
										VALUES(?,?,?,?,?,?)`

	result, err := repo.db.Exec(sql, params.ProductCode, params.ProductName,
		params.Category, params.Description, params.ListPrice, params.StandardCost)
	helper.Catch(err)

	id, _ := result.LastInsertId()

	return id, nil
}

func (repo *repository) UpdateProducts(params *updateProductRequest) (int64, error) {
	const sql = `UPDATE products 
					SET 
					product_code = ?, 
					product_name = ?,
					category = ?,
					description = ?,
					list_price = ?,
					standard_cost = ?
					WHERE id = ?
				`

	_, err := repo.db.Exec(sql, params.ProductCode, params.ProductName,
		params.Category, params.Description, params.ListPrice, params.StandardCost, params.ID)
	helper.Catch(err)

	return params.ID, nil
}

func (repo *repository) DeleteProducts(params *deleteProductRequest) (int64, error) {

	const sql = `DELETE  FROM products WHERE id = ?`

	result, err := repo.db.Exec(sql, params.ProductID)

	helper.Catch(err)
	count, err := result.RowsAffected()
	helper.Catch(err)

	return count, nil
}

func (repo *repository) getBestSellers() ([]*ProductTop, error) {

	const sql = `SELECT 
						od.product_id,
						p.product_name,
						SUM(od.quantity*od.unit_price) vendido
						FROM order_details od
						inner 
						join products p on od.product_id = p.id
						GROUP by od.product_id
						ORDER BY vendido desc 
						limit 10`

	result, err := repo.db.Query(sql)

	helper.Catch(err)
	var products []*ProductTop

	for result.Next() {
		product := &ProductTop{}
		err = result.Scan(&product.ID, &product.ProductName, &product.Vendidos)
		helper.Catch(err)

		products = append(products, product)
	}

	return products, nil

}

func (repo *repository) getTotalVentas() (float64, error) {
	const sql = `SELECT SUM(od.quantity*od.unit_price) vendido FROM order_details od`
	var total float64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}
