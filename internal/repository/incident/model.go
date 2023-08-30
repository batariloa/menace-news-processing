package incident

import "time"

type StoredIncident struct {
	Id        int       `json:id bson:id`
	Title     string    `json:"title" bson:"title"`
	Latitude  float64   `json:"latitude" bson:"latitude"`
	Longitude float64   `json:"longitude" bson:"longitude"`
	Date      time.Time `json:"date" bson:"date"`
}
