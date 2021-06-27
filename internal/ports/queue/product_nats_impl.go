package queue

import (
	"fmt"
	"github.com/beuus39/product/internal/shared/dtos"
	"github.com/beuus39/product/pkg/nats"
)

type ProductQueue struct {
	nats nats.NatsConfig
}

func (p *ProductQueue) PublishProduct(topic string, product *dtos.ProductDto) bool {
	c, err := p.nats.Connect()
	defer c.Close()

	if err != nil {
		 return false
	}
	err = c.Publish(topic, product)
	if err != nil {
		return false
	}
	return true
}

func (p *ProductQueue) SubscriberHello(topic string) {
	c, err := p.nats.Connect()
	defer c.Close()

	if err != nil {
		return
	}
	c.Subscribe(topic, func(s string) {
		fmt.Printf("Received a message: %s\n", s)
	})
}

func (p *ProductQueue) SendSayHello(topic string, body interface{}) bool {
	c, err := p.nats.Connect()
	defer c.Close()
	if err != nil {
		return false
	}
	err = c.Publish(topic, body)
	if err != nil {
		return false
	}
	return true
}

func NewProductQueue(nats nats.NatsConfig) ProductAmqp {
	return &ProductQueue{
		nats: nats,
	}
}
