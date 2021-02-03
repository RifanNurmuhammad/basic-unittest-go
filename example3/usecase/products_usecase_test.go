package usecase

import (
	"basic-unittest-go/example3/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductsRepositoryMock{Mock: mock.Mock{}}
var producUseCase = ProductsUseCase{Repository: productRepository}

func TestProductsUseCase_Get(t *testing.T) {
	productRepository.Mock.On("FindByID", "1").Return(nil)
	product, err := producUseCase.Get("1")
	assert.Nil(t, product)
	assert.NotNil(t, err)
}
