package main

import (
	"GolangValdezApi/databasecon"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// "net/http"
// "database/sql"
// "github.com/go-chi/chi"
// "github.com/go-chi/chi/middleware"

func main() {

	dbco := databasecon.InitDB()
	defer dbco.Close()
	fmt.Println(dbco)

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// http.ListenAndServe(":3000", r)
}
