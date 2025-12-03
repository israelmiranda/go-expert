package tax

import "testing"

func FuzzCalculateTax(f *testing.F) {
	// Arrange
	seed := []float64{-1, 0, 1, 1000, 10000.0}
	for _, amount := range seed {
		f.Add(amount)
	}

	// Act & Assert
	f.Fuzz(func(t *testing.T, amount float64) {
		// Act
		result := CalculateTax(amount)
		if result == 0 && amount != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
		if result == 10.0 && amount >= 1000 {
			t.Errorf("Received %f but expected 0", result)
		}
		if result == 20.0 && amount >= 20000 {
			t.Errorf("Received %f but expected 0", result)
		}
	})
}
