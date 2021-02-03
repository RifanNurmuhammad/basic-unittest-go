package repository

import "basic-unittest-go/example3/domain"

// ProductsRepository interface
type ProductsRepository interface {
	FindByID(id string) *domain.Product
}
