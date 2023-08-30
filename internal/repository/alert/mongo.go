package alert

import (
	"context"
	"time"

	"github.com/batariloa/bobber/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

const alerts_collection = "incidents"

type AlertsMongoStore struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewAlertsMongoStore(client *mongo.Client) *AlertsMongoStore {

	return &AlertsMongoStore{
		client:     client,
		database:   repository.Database,
		collection: alerts_collection,
	}
}

func (s *AlertsMongoStore) Save(alert *Alert) error {

	coll := s.client.Database(s.database).Collection(s.collection)

	_, err := coll.InsertOne(context.TODO(), *alert)

	if err != nil {
		return err
	}

	return nil
}

func (s *AlertsMongoStore) GetLatestAlertDate() (time.Time, error) {

	return time.Now(), nil
}
