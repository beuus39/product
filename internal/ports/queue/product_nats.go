package queue

import "github.com/beuus39/product/internal/shared/dtos"

type ProductAmqp interface {
	SendSayHello(topic string, body interface{}) bool
	SubscriberHello(topic string)
	PublishProduct(topic string, product *dtos.ProductDto) bool
}
