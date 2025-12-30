//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/israelmiranda/go-expert/di/2/infra"
	"github.com/israelmiranda/go-expert/di/2/product"
)

var addRepositories = wire.NewSet(
	infra.NewProductRepository,
	wire.Bind(new(product.ProductRepository), new(*infra.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		addRepositories,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
