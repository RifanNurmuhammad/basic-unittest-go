package usecase

import (
	"fmt"

	"basic-unittest-go/example3/domain"
	"basic-unittest-go/example3/repository"
)

// ProductsUseCase struct
type ProductsUseCase struct {
	Repository repository.ProductsRepository
}

// Get functions for getting the product
func (usecase ProductsUseCase) Get(id string) (*domain.Product, error) {
	product := usecase.Repository.FindByID(id)
	if product == nil {
		return nil, fmt.Errorf("cannot find product %s", id)
	}
	return product, nil
}
