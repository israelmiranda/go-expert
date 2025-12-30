package product

type ProductSaver interface {
	SaveProduct(product Product) error
}

type SaveProductUseCase struct {
	repository ProductSaver
}

func NewSaveProductUseCase(repository ProductSaver) *SaveProductUseCase {
	return &SaveProductUseCase{repository}
}

func (u *SaveProductUseCase) SaveProduct(product Product) error {
	return u.repository.SaveProduct(product)
}
