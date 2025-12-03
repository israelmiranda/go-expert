package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	// Arrange
	amount := 1000.0
	expected := 10.0

	// Act
	result, err := CalculateTax(amount)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestCalculateTaxWithError(t *testing.T) {
	// Act
	result, err := CalculateTax(0)

	// Assert
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, result)
}
