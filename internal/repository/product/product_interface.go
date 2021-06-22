package product

import (
	"github.com/beuus39/product/internal/domain"
)

type RepositoryResult struct {
	Result interface{}
	Error error
}
type Repository interface {
	Find(id int) <-chan RepositoryResult
	Save(product *domain.Product) <-chan RepositoryResult
	Update(product *domain.Product) <-chan RepositoryResult
	FindAll() <-chan RepositoryResult
	Delete(id int) <-chan RepositoryResult
}