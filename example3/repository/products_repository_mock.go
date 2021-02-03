package repository

import (
	"basic-unittest-go/example3/domain"

	"github.com/stretchr/testify/mock"
)

// ProductsRepositoryMock struct
type ProductsRepositoryMock struct {
	Mock mock.Mock
}

// FindByID mock function
func (repository *ProductsRepositoryMock) FindByID(id string) *domain.Product {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(domain.Product)
	return &product
}
