package app

import (
	"github.com/beuus39/product/internal/domain"
	"github.com/beuus39/product/internal/repository/product"
	"github.com/pkg/errors"
)

type service struct {
	repo product.Repository
}

func (s *service) FindProduct(id int) <-chan ServiceResult {
	output := make(chan ServiceResult)

	go func() {
		productResult := <-s.repo.Find(id)
		if productResult.Error != nil {
			output <- ServiceResult{Error: productResult.Error}
			return
		}

		product, ok := productResult.Result.(*domain.Product)
		if !ok {
			output <- ServiceResult{Error: errors.New("result is not product")}
			return
		}

		output <- ServiceResult{Result: product}
	}()
	return output
}

func (s *service) SaveProduct(product *domain.Product) <-chan ServiceResult {
	output := make(chan ServiceResult)

	go func() {
		productResult := <-s.repo.Save(product)

		if productResult.Error != nil {
			output <- ServiceResult{Error: productResult.Error}
		}

		output <- ServiceResult{ Result: productResult.Result}
	}()
	return output
}

func (s *service) UpdateProduct(product *domain.Product) <-chan ServiceResult {
	output := make(chan ServiceResult)

	go func() {
		productResult := <-s.repo.Update(product)

		if productResult.Error != nil {
			output <- ServiceResult{Error: productResult.Error}
		}

		output <- ServiceResult{ Result: productResult.Result}
	}()
	return output
}

func (s *service) FindAllProducts() <-chan ServiceResult {
	output := make(chan ServiceResult)

	go func() {
		productResult := <-s.repo.FindAll()

		if productResult.Error != nil {
			output <- ServiceResult{Error: productResult.Error}
		}

		products, ok := productResult.Result.([]*domain.Product)
		if !ok {
			output <- ServiceResult{Error: errors.New("result is not product")}
		}
		output <- ServiceResult{ Result: products}
	}()
	return output
}

func (s *service) DeleteProduct(id int) <-chan ServiceResult {
	output := make(chan ServiceResult)

	go func() {
		productResult := <-s.repo.Delete(id)

		if productResult.Error != nil {
			output <- ServiceResult{Error: productResult.Error}
		}

		output <- ServiceResult{ Result: productResult.Result}
	}()
	return output
}

func NewProductService(repo product.Repository) Service {
	return &service{
		repo: repo,
	}
}