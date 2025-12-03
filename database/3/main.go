package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model // base model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// create category
	// category1 := Category{
	// 	Name: "category 1",
	// }
	// db.Create(&category1)

	// category2 := Category{
	// 	Name: "category 2",
	// }
	// db.Create(&category2)

	// create product
	// db.Create(&Product{
	// 	Name:       "product 1",
	// 	Price:      100.1,
	// 	CategoryID: category.ID,
	// })

	// create product batch
	// products := []Product{
	// 	{Name: "product 1", Price: 100.1, CategoryID: category1.ID},
	// 	{Name: "product 2", Price: 100.2, CategoryID: category1.ID},
	// 	{Name: "product 3", Price: 100.3, CategoryID: category2.ID},
	// 	{Name: "product 4", Price: 100.4, CategoryID: category2.ID},
	// }
	// db.Create(products)

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Printf("%s - %s\n", product.Category.Name, product.Name)
	}
}
