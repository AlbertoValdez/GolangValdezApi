package databasecon

import (
	"database/sql"
	"fmt"

	//Only si se usa jaja
	_ "github.com/go-sql-driver/mysql"
)

//InitDB SQL
func InitDB() *sql.DB {

	conectionurl := "root:udl131296.@tcp(127.0.0.1:3306)/northwind"
	databasec, err := sql.Open("mysql", conectionurl)

	if err != nil {
		fmt.Println("error en el conection", err)
		panic(err.Error()) //error Handling
	}

	return databasec

}
