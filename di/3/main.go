package main

import (
	"database/sql"
	"fmt"

	"github.com/israelmiranda/go-expert/di/3/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	saveUseCase := NewSaveProductUseCase(db)
	err = saveUseCase.SaveProduct(product.Product{
		ID:   1,
		Name: "Product Name",
	})
	if err != nil {
		panic(err)
	}

	getUseCase := NewGetProductUseCase(db)
	product, err := getUseCase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
