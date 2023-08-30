package subscriber

import (
	"context"

	"github.com/batariloa/bobber/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubscriberMongoStore struct {
	client     *mongo.Client
	database   string
	collection string
}

const collection = "subscriber"

func NewSubscriberMongoStore(client *mongo.Client) *SubscriberMongoStore {

	return &SubscriberMongoStore{
		client:     client,
		database:   repository.Database,
		collection: collection,
	}
}
func (s *SubscriberMongoStore) GetAllSubscribers() ([]Subscriber, error) {
	coll := s.client.Database(s.database).Collection(s.collection)

	// Define options for limiting the number of results to 50
	options := options.Find()
	options.SetLimit(50)

	// Find the first 50 subscribers
	cursor, err := coll.Find(context.TODO(), bson.D{}, options)
	if err != nil {
		return nil, err
	}

	results := []Subscriber{}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *SubscriberMongoStore) Subscribe(subscriber *Subscriber) error {

	coll := s.client.Database(s.database).Collection(s.collection)

	_, err := coll.InsertOne(context.TODO(), subscriber)
	if err != nil {
		return err
	}
	return nil
}
