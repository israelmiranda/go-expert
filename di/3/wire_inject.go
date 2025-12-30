//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/israelmiranda/go-expert/di/3/infra"
	"github.com/israelmiranda/go-expert/di/3/product"
)

var repositoryProvider = wire.NewSet(infra.NewProductRepository)

var getProductSet = wire.NewSet(
	repositoryProvider,
	wire.Bind(new(product.ProductGetter), new(*infra.ProductRepository)),
)

var saveProductSet = wire.NewSet(
	repositoryProvider,
	wire.Bind(new(product.ProductSaver), new(*infra.ProductRepository)),
)

func NewGetProductUseCase(db *sql.DB) *product.GetProductUseCase {
	wire.Build(
		getProductSet,
		product.NewGetProductUseCase,
	)
	return &product.GetProductUseCase{}
}

func NewSaveProductUseCase(db *sql.DB) *product.SaveProductUseCase {
	wire.Build(
		saveProductSet,
		product.NewSaveProductUseCase,
	)
	return &product.SaveProductUseCase{}
}
