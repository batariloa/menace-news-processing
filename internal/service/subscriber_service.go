package service

import (
	"github.com/batariloa/bobber/internal/repository/subscriber"
)

type SubscriberStore = subscriber.SubscriberStore

type SubscriberService struct {
	store subscriber.SubscriberStore
}

func NewSubscriberService(store SubscriberStore) *SubscriberService {

	return &SubscriberService{
		store: store,
	}
}

func (s *SubscriberService) SaveSubscriber(subscriber *subscriber.Subscriber) error {
	err := s.store.Subscribe(subscriber)
	if err != nil {
		return err
	}

	return nil
}

func (s *SubscriberService) GetAllSubscribers() ([]subscriber.Subscriber, error) {

	return s.store.GetAllSubscribers()
}
