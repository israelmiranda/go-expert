package product

type ProductRepository interface {
	GetProduct(id int) (Product, error)
}
