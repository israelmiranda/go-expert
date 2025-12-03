package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	gorm.Model // base model
}

func main() {
	// dsn := "root:root@tcp(localhost:3306)/goexpert"
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	// db.Debug().Create(&Product{
	// 	Name:  "product 1",
	// 	Price: 100.2,
	// })

	// create batch
	// products := []Product{
	// 	{Name: "product 1", Price: 100.1},
	// 	{Name: "product 2", Price: 100.2},
	// 	{Name: "product 3", Price: 100.3},
	// }
	// db.Debug().Create(products)

	// select one
	// var product Product
	// db.Debug().First(&product, 1)
	// fmt.Println(product)

	// db.Debug().First(&product, "name = ?", "product 2")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Debug().Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Debug().Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Debug().Where("price > ?", 100.1).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// like
	// var products []Product
	// db.Debug().Where("name LIKE ?", "%2%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// update & delete
	// var product1 Product
	// db.Debug().First(&product1, 1)
	// product1.Name = "product 4"
	// db.Debug().Save(&product1)

	// var product2 Product
	// db.Debug().First(&product2, 1)
	// fmt.Println(product2)

	// db.Debug().Delete(&product2)
}
