package amqp

import (
	"github.com/streadway/amqp"
	"log"
)

type AmqpConfig struct {
	Uri string
	Queue string
	Channel string
}

func (a *AmqpConfig) GetChanel() (ch *amqp.Channel, err error) {
	conn, err := amqp.Dial(a.Uri)
	defer conn.Close()
	if err != nil {
		log.Fatal("Failed to connect RabbitMq")
	}
	ch, err = conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Fatal("Failed to connect to channel")
	}
	_, err = ch.QueueDeclare(
		a.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	return ch, err
}

func (a *AmqpConfig) Publish(contentType string, body string) (err error) {
	ch, err := a.GetChanel()
	if err != nil {
		log.Fatal("Failed to connect channel")
	}
	err = ch.Publish(
		"product.queue",
		"hello",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: contentType,
			Body: []byte(body),
		})
	return err
}

func (a *AmqpConfig) Subscribe() {
	ch, err := a.GetChanel()
	defer ch.Close()
	if err != nil {
		log.Fatal("Failed to get channel")
	}
	msgs, err := ch.Consume(
		a.Queue,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatal("Failed to register a consumer")
	}
	forever := make(chan []byte)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	<-forever
}

func NewAmqpConnector(conf *AmqpConfig) (conn AmqpConnector) {
	return &AmqpConfig{
		Uri: conf.Uri,
		Queue: conf.Queue,
		Channel: conf.Channel,
	}
}
