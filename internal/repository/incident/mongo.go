package incident

import (
	"context"
	"fmt"

	"github.com/batariloa/bobber/internal/repository"
	"github.com/batariloa/bobber/internal/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const incidents_collection = "incidents"

type IncidentMongoStore struct {
	client     *mongo.Client
	database   string
	collection string
}

func (store *IncidentMongoStore) GetAllFromToday() ([]*StoredIncident, error) {
	coll := store.client.Database(store.database).Collection(store.collection)

	query := util.GetIsTodayQuery()

	cursor, err := coll.Find(context.TODO(), query)

	if err != nil {
		return nil, err
	}

	results := []*StoredIncident{}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (store *IncidentMongoStore) GetAll() ([]*StoredIncident, error) {
	coll := store.client.Database(store.database).Collection(store.collection)

	query := bson.M{}

	cursor, err := coll.Find(context.TODO(), query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	results := []*StoredIncident{}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (store *IncidentMongoStore) Save(i *StoredIncident) error {
	coll := store.client.Database(store.database).Collection(store.collection)

	_, err := coll.InsertOne(context.TODO(), i)
	return err
}

func (store *IncidentMongoStore) SaveAll(incidents []*StoredIncident) error {
	// Convert the slice of incidents to a slice of interface{}
	documents := make([]interface{}, len(incidents))
	for i, dto := range incidents {
		documents[i] = dto
	}

	// Insert the documents in a single database call
	coll := store.client.Database(store.database).Collection(store.collection)
	_, err := coll.InsertMany(context.TODO(), documents)
	return err
}

func NewIncidentMongoStore(client *mongo.Client) *IncidentMongoStore {

	return &IncidentMongoStore{
		client:     client,
		database:   repository.Database,
		collection: incidents_collection,
	}
}
