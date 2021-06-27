package app

import (
	"github.com/beuus39/product/internal/domain"
)

type ServiceResult struct {
	Result interface{}
	Error error
}

type Service interface {
	FindProduct(id int) <-chan ServiceResult
	SaveProduct(product *domain.Product) <-chan ServiceResult
	UpdateProduct(product *domain.Product) <-chan ServiceResult
	FindAllProducts() <-chan ServiceResult
	DeleteProduct(id int) <-chan ServiceResult
}