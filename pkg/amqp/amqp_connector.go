package amqp

import "github.com/streadway/amqp"

type AmqpConnector interface {
	GetChanel() (ch *amqp.Channel, err error)
	Publish(contentType string, body string) (err error)
	Subscribe()
}
