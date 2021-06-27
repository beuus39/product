package adapters

import (
	"fmt"
	"github.com/beuus39/product/internal/shared/dtos"
	"github.com/beuus39/product/pkg/nats"
)

type OrderNatsConfiguration struct {
	nats nats.NatsConnector
}

func (o *OrderNatsConfiguration) SubscriberOrder(topic string) {
	c, err := o.nats.Connect()
	defer c.Close()

	if err != nil {
		return
	}
	_, err = c.Subscribe(topic, func(subj, reply string, order *dtos.OrderSubscriber) {
		fmt.Printf("Received a product on subject %s! %+v\n", subj, order)
	})
	if err != nil {
		return
	}
}

func NewOrderSubscriber(nats nats.NatsConnector) OrderSubscriber {
	return &OrderNatsConfiguration{
		nats: nats,
	}
}
