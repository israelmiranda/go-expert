package tax

type Repository interface {
	Save(tax float64) error
}

func CalculateTaxAndSave(repository Repository, amount float64) error {
	tax := CalculateTax(amount)
	return repository.Save(tax)
}

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
