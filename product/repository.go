package product

import "database/sql"

//Repository interface
type Repository interface {
	GetProductByID(productID int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProducts(params *getAddProductRequest) (int64, error)
	UpdateProducts(params *updateProductRequest) (int64, error)
	DeleteProducts(params *deleteProductRequest) (int64, error)
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
	if err != nil {
		panic(err)
	}
	return total, nil
}

func (repo *repository) InsertProducts(params *getAddProductRequest) (int64, error) {
	const sql = `INSERT INTO products (product_code, product_name,category,description,list_price,standard_cost)
										VALUES(?,?,?,?,?,?)`

	result, err := repo.db.Exec(sql, params.ProductCode, params.ProductName,
		params.Category, params.Description, params.ListPrice, params.StandardCost)
	if err != nil {

		panic(err)

	}

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
	if err != nil {

		panic(err)

	}

	return params.ID, nil
}

func (repo *repository) DeleteProducts(params *deleteProductRequest) (int64, error) {

	const sql = `DELETE  FROM products WHERE id = ?`

	result, err := repo.db.Exec(sql, params.ProductID)

	if err != nil {
		panic(err)

	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)

	}

	return count, nil
}
