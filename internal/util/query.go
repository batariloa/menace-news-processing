package util

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetIsTodayQuery() bson.M {

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	return bson.M{
		"date": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}
}
