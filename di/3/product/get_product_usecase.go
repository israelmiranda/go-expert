package product

type ProductGetter interface {
	GetProduct(id int) (Product, error)
}

type GetProductUseCase struct {
	repository ProductGetter
}

func NewGetProductUseCase(repository ProductGetter) *GetProductUseCase {
	return &GetProductUseCase{repository}
}

func (u *GetProductUseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}
