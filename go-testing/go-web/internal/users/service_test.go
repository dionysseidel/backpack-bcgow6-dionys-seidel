package users

import (
	"errors"
	"fmt"
	"testing"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/internal/domain"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationDelete(t *testing.T) {
	// Arrange
	initialDatabase := []domain.User{{
		ID:       1,
		Name:     "Dionys",
		IsActive: true,
		Age:      25,
	}, {
		ID:       2,
		Name:     "Brian",
		IsActive: true,
		Age:      28,
	}, {
		ID:       3,
		Name:     "Matías",
		IsActive: true,
		Age:      25,
	}}

	iDToDelete := initialDatabase[2].ID

	reducedDatabase := []domain.User{{
		ID:       1,
		Name:     "Dionys",
		IsActive: true,
		Age:      25,
	}, {
		ID:       2,
		Name:     "Brian",
		IsActive: true,
		Age:      28,
	}}

	mockStorage := &mock.MockStorage{
		DataMock:       initialDatabase,
		ReadWasCalled:  false,
		WriteWasCalled: false,
	}
	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	err := service.Delete(iDToDelete)

	// Assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.WriteWasCalled)
	assert.Equal(t, reducedDatabase, mockStorage.DataMock)
	assert.True(t, mockStorage.ReadWasCalled)
}

func TestServiceIntegrationDeleteFail(t *testing.T) {
	// Arrange
	initialDatabase := []domain.User{{
		ID:       1,
		Name:     "Dionys",
		IsActive: true,
		Age:      25,
	}, {
		ID:       2,
		Name:     "Brian",
		IsActive: true,
		Age:      28,
	}, {
		ID:       3,
		Name:     "Matías",
		IsActive: true,
		Age:      25,
	}}

	nonExistingID := initialDatabase[2].ID

	mockStorage := &mock.MockStorage{
		DataMock:       initialDatabase,
		ReadWasCalled:  false,
		WriteWasCalled: false,
	}
	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	err := service.Delete(nonExistingID)

	// Assert
	assert.NotNil(t, err)
	assert.True(t, mockStorage.ReadWasCalled)
	assert.False(t, mockStorage.WriteWasCalled)
}

func TestServiceIntegrationGetAll(t *testing.T) {
	// Arrange
	database := []domain.User{{
		ID:       1,
		Name:     "Dionys",
		IsActive: true,
		Age:      25,
	}, {
		ID:       2,
		Name:     "Brian",
		IsActive: true,
		Age:      28,
	}}

	mockStorage := &mock.MockStorage{
		DataMock: database,
	}
	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, database, result)
}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	// Arrange
	expectedError := errors.New("can't read database")

	mockStorage := &mock.MockStorage{
		DataMock:   nil,
		ErrorWrite: nil,
		ErrorRead:  expectedError,
	}
	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.GetAll()

	// Assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, result)
}

func TestServiceIntegrationStore(t *testing.T) {
	// Arrange
	initialDatabase := []domain.User{{
		ID:       1,
		Name:     "Dionys",
		IsActive: true,
		Age:      25,
	}, {
		ID:       2,
		Name:     "Brian",
		IsActive: true,
		Age:      28,
	}}

	newUser := domain.User{
		ID:       3,
		Name:     "Matías",
		IsActive: true,
		Age:      25,
	}

	expected := append(initialDatabase, newUser)

	mockStorage := &mock.MockStorage{
		DataMock:   initialDatabase,
		ErrorRead:  nil,
		ErrorWrite: nil,
	}
	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.Store(
		newUser.Name,
		newUser.IsActive,
		newUser.Age,
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expected, mockStorage.DataMock)
	assert.Equal(t, result, newUser)
	assert.Equal(t, newUser.ID, mockStorage.DataMock[len(mockStorage.DataMock)-1].ID)
}

func TestServiceIntegrationStoreFail(t *testing.T) {
	// Arrange
	newUser := domain.User{
		ID:       3,
		Name:     "Matías",
		IsActive: true,
		Age:      25,
	}

	writeError := fmt.Errorf("can't write in database")

	expectedError := fmt.Errorf("error creating user: %w", writeError)

	mockStorage := &mock.MockStorage{
		DataMock:   nil,
		ErrorRead:  nil,
		ErrorWrite: expectedError,
	}
	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.Store(newUser.Name, newUser.IsActive, newUser.Age)
	// result.ID = 12345678 // Debería lanzar error

	// Assert
	assert.EqualError(t, expectedError, err.Error())
	assert.Equal(t, domain.User{}, result)
	assert.Empty(t, result)
}

func TestServiceIntegrationUpdate(t *testing.T) {
	// Arrange
	initialDatabase := []domain.User{{
		ID:       1,
		Name:     "Dionys",
		IsActive: true,
		Age:      25,
	}, {
		ID:       2,
		Name:     "Brian",
		IsActive: true,
		Age:      28,
	}, {
		ID:       3,
		Name:     "Matías",
		IsActive: true,
		Age:      25,
	}}

	userToUpdate := domain.User{
		ID:       3,
		Name:     "Matías",
		IsActive: false,
		Age:      25,
	}

	mockStorage := &mock.MockStorage{
		DataMock:       initialDatabase,
		ReadWasCalled:  false,
		WriteWasCalled: false,
	}
	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.Update(userToUpdate.ID, userToUpdate.Name, userToUpdate.IsActive, userToUpdate.Age)

	// Assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.WriteWasCalled)
	assert.Equal(t, userToUpdate.ID, mockStorage.DataMock[userToUpdate.ID-1].ID)
	assert.Equal(t, userToUpdate, result)
	assert.True(t, mockStorage.ReadWasCalled)
}
