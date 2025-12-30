package main

import (
	"database/sql"
	"fmt"

	"github.com/israelmiranda/go-expert/di/2/infra"
	"github.com/israelmiranda/go-expert/di/2/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := infra.NewProductRepository(db)
	usecase := product.NewProductUseCase(repository)

	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
