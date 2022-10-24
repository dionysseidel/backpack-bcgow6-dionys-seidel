package mock

import (
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/internal/domain"
)

type MockStorage struct {
	DataMock       []domain.User
	ErrorWrite     error
	ErrorRead      error
	ReadWasCalled  bool
	WriteWasCalled bool
}

func (ms *MockStorage) Read(data interface{}) error {
	ms.ReadWasCalled = true

	if ms.ErrorRead != nil {
		return ms.ErrorRead //fmt.Errorf(ms.ErrorRead)
	}

	castedData := data.(*[]domain.User) // El puntero es necesario para sobreescribir data
	*castedData = ms.DataMock

	return nil
}

func (ms *MockStorage) Write(data interface{}) error {
	ms.WriteWasCalled = true

	if ms.ErrorWrite != nil {
		return ms.ErrorWrite //fmt.Errorf(ms.ErrorWrite)
	}

	castedData := data.([]domain.User)
	// ms.DataMock = append(ms.DataMock, castedData[len(castedData)-1])
	ms.DataMock = castedData

	return nil
}
