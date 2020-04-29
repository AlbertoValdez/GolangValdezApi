package product

import "database/sql"

//Repository interface
type Repository interface {
	GetProductByID(productID int) (*Product, error)
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
	if err != nil {
		panic(err)
	}

	return product, err
}
