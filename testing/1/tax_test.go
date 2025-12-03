package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	// Arrange
	amount := 500.0
	expected := 5.0

	// Act
	result := CalculateTax(amount)

	// Assert
	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCaculateTaxBatch(t *testing.T) {
	// Arrange
	type Tax struct {
		amount, expected float64
	}
	table := []Tax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		// {0.0, 0.0},
	}

	// Act & Assert
	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, item.amount)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	// for i := 0; i < b.N; i++ { deprecated
	for b.Loop() {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for b.Loop() {
		CalculateTax2(500.0)
	}
}
