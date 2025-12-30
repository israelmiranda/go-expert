package infra

import (
	"database/sql"
	"fmt"

	"github.com/israelmiranda/go-expert/di/3/product"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) GetProduct(id int) (product.Product, error) {
	return product.Product{
		ID:   id,
		Name: "Product Name",
	}, nil
}

func (r *ProductRepository) SaveProduct(product product.Product) error {
	fmt.Println("product saved", product)
	return nil
}
