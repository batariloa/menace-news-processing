package service

import (
	"math/rand"
	"time"

	"github.com/batariloa/bobber/internal/datastruct"
	"github.com/batariloa/bobber/internal/repository/incident"
)

func GenerateStoredIncidents(fetched *[]datastruct.FetchedIncident) []*incident.StoredIncident {

	sl := []*incident.StoredIncident{}

	for _, i := range *fetched {

		sl = append(sl, &incident.StoredIncident{
			Title:     i.Title,
			Latitude:  generateLatitude(),
			Longitude: generateLongitude(),
			Date:      time.Now(),
		})
	}

	return sl
}

func generateLatitude() float64 {
	// Generate a random number between 0 and 1
	random := rand.Float64()
	// Scale the random number to the desired range (46.05 - 40) and add the minimum value (40)
	number := random*(46-42) + 42
	return number
}

func generateLongitude() float64 {
	// Generate a random number between 0 and 1
	random := rand.Float64()
	// Scale the random number to the desired range (46.05 - 40) and add the minimum value (40)
	number := random*(25-21) + 21

	return number
}
