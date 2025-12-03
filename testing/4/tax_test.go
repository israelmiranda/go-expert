package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCalculateTaxAndSaveSuccessfully(t *testing.T) {
	// Arrange
	repository := &RepositoryMock{}
	repository.On("Save", mock.Anything).Return(nil)

	// Act
	err := CalculateTaxAndSave(repository, 1000)

	// Assert
	assert.Nil(t, err)

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "Save", 1)
}

func TestCalculateTaxAndSaveFail(t *testing.T) {
	// Arrange
	repository := &RepositoryMock{}
	repository.On("Save", 0.0).Return(errors.New("tax must be greater than zero"))

	// Act
	err := CalculateTaxAndSave(repository, 0.0)

	// Assert
	assert.Error(t, err, "tax must be greater than zero")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "Save", 1)
}
