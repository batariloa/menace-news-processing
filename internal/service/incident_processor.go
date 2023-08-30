package service

import (
	"regexp"
	"strings"

	"github.com/batariloa/bobber/internal/datastruct"
)

type IncidentService struct{}

func NewIncidentService() *IncidentService {
	return &IncidentService{}
}

func (i *IncidentService) MoldIncidents(incidents []datastruct.FetchedIncident) []datastruct.FetchedIncident {
	var phrases = []string{"Florida Man", "Crazy man", "Biden"}

	incidentDoerName := "Velja Nevolja"

	result := i.filterMatches(incidents, phrases, incidentDoerName)

	return result
}

func (*IncidentService) filterMatches(incidents []datastruct.FetchedIncident, patterns []string, replacement string) []datastruct.FetchedIncident {
	var modifiedIncidents []datastruct.FetchedIncident

	// Create a regular expression pattern that matches any of the given patterns
	patternRegex := strings.Join(patterns, "|")
	re := regexp.MustCompile(patternRegex)

	for _, i := range incidents {
		modifiedTitle := re.ReplaceAllString(i.Title, replacement)

		// Add the datastruct to the modifiedIncidents slice if the title was modified
		if modifiedTitle != i.Title {
			modifiedIncidents = append(modifiedIncidents, datastruct.FetchedIncident{Title: modifiedTitle})
		}
	}

	return modifiedIncidents
}
