package infra

import (
	"database/sql"

	"github.com/israelmiranda/go-expert/di/2/product"
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
