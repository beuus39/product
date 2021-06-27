package dtos

import (
	"github.com/beuus39/product/internal/domain"
	"github.com/beuus39/product/internal/shared/dtos"
)

func OrderToProductConverter(subscriber dtos.OrderSubscriber) (product domain.Product)  {
	product.ID = subscriber.ID
	return product
}
