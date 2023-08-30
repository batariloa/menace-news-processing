package alert

import (
	"time"
)

type Alert struct {
	Time       time.Time `json:"date" bson:"date"`
	IncidentID string    `json:"incidentId" bson:"incidentId"`
}
