package costumers

import (
	"database/sql"

	"github.com/GolangValdezApi/helper"
)

type Repository interface {
	getClients(params *getCostumerListRequest) ([]*Customer, error)
	GetTotalClients() (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {

	return &repository{
		db: db,
	}
}

func (repo *repository) getClients(params *getCostumerListRequest) ([]*Customer, error) {

	const sql = `SELECT 
				c.id,
				c.first_name,
				c.last_name,
				address,
				business_phone,
				city,
				company

				FROM customers c
				LIMIT ? OFFSET ?`
	result, err := repo.db.Query(sql, params.limit, params.offset)
	helper.Catch(err)
	var clientes []*Customer

	for result.Next() {
		cliente := &Customer{}
		err = result.Scan(&cliente.ID, &cliente.FirtsName, &cliente.LastName,
			&cliente.Address, &cliente.BusniessPhone,
			&cliente.City, &cliente.Company)
		helper.Catch(err)
		clientes = append(clientes, cliente)
	}
	return clientes, nil

}

func (repo *repository) GetTotalClients() (int64, error) {

	const sql = `SELECT COUNT(*) FROM customers`
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil

}
