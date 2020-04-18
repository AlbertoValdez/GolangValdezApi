package databasecon

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {

	conectionurl := "root:root@tcp(localhost:3306)/northwind"
	databasec, err := sql.Open("mysql", conectionurl)

	if err != nil {
		panic(err.Error()) //error Handling manejo de errores
	}

	return databasec

}
