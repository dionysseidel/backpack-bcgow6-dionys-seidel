package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type RepositoryMock struct {
	getAllBySellerWasCalled bool
}

func (r *RepositoryMock) GetAllBySeller(sellerID string) ([]Product, error) {
	r.getAllBySellerWasCalled = true
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	return prodList, nil
}

func TestGetAllBySeller(t *testing.T) {
	productExpected := Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	}

	repositoryMock := &RepositoryMock{
		getAllBySellerWasCalled: false,
	}
	service := NewService(repositoryMock)

	products, err := service.GetAllBySeller("mock")

	assert.Nil(t, err)
	assert.True(t, repositoryMock.getAllBySellerWasCalled)
	assert.Contains(t, products, productExpected)
}
