package users

import (
	"fmt"
	"testing"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/internal/domain"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/test/mock"
	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock   []domain.User
	errorWrite string
	errorRead  string
}

func (ms *MockStorage) Read(data interface{}) error {
	if ms.errorRead != "" {
		return fmt.Errorf(ms.errorRead)
	}
	castedData := data.(*[]domain.User) // El puntero es necesario para sobreescribir data
	*castedData = ms.dataMock
	return nil
}

func (ms *MockStorage) Write(data interface{}) error {
	if ms.errorWrite != "" {
		return fmt.Errorf(ms.errorWrite)
	}
	castedData := data.([]domain.User)
	ms.dataMock = append(ms.dataMock, castedData[len(castedData)-1])
	return nil
}

type StubRepository struct{}

func /*(sr StubRepository)*/ TestGetAll(t *testing.T) {
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
		// ErrorWrite: "",
		// ErrorRead:  "",
	}

	repository := NewRepository(mockStorage)

	// Act
	// mockStorage.DataMock = append(mockStorage.DataMock, domain.User{}) // Para fallar la prueba
	result, err := repository.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, database, result)
}

// func TestFindByName(t *testing.T) {
// 	//arrange

// 	stubDB := StubRepository{}
// 	motor := NewRepository()
// 	telefonoEsperado := "12345678"

// 	//act
// 	resultado := motor.FindByName("Nacho")

// 	//assert
// 	assert.Equal(t, telefonoEsperado, resultado)
// }
