package subscriber

type SubscriberStore interface {
	Subscribe(*Subscriber) error
	GetAllSubscribers() ([]Subscriber, error)
}
