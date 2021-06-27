package adapters

type OrderSubscriber interface {
	SubscriberOrder(topic string)
}
