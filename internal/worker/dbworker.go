package worker

import (
	"fmt"
	"sync"

	"github.com/batariloa/bobber/internal/datastruct"
	"github.com/batariloa/bobber/internal/repository/incident"
	"github.com/batariloa/bobber/internal/service"
)

type fi = datastruct.FetchedIncident
type AlertService = service.AlertService
type IncidentStorer = incident.IncidentStorer

type StorerWorker struct {
	incidentCh   *chan *[]fi
	storer       IncidentStorer
	mutex        *sync.Mutex
	alertService AlertService
}

func NewStorerWorker(incidentCh *chan *[]fi, store IncidentStorer, mutex *sync.Mutex, as *AlertService) *StorerWorker {
	return &StorerWorker{
		incidentCh:   incidentCh,
		storer:       store,
		mutex:        mutex,
		alertService: *as,
	}
}

func (sw *StorerWorker) Start() {

	fmt.Println("Started SW.")

	unique, err := sw.buildIncidentHashset()
	if err != nil {
		panic("Unhandled database error.")
	}

	for incidents := range *sw.incidentCh {

		fmt.Println("Receiving incidents.")

		sw.mutex.Lock()
		filtered := filterNovelIncidents(*incidents, &unique)

		generatedIncidents := service.GenerateStoredIncidents(&filtered)
		sw.storer.SaveAll(generatedIncidents)

		sw.mutex.Unlock()

		fmt.Println("Finished storing incidents. Number of new incidents: ", len(filtered))

		if len(generatedIncidents) > 0 {
			sw.alertService.AlertUsers(generatedIncidents[0])
		}
	}
}

func filterNovelIncidents(incidents []fi, unique *map[string]bool) []fi {

	filtered := []fi{}

	for _, item := range incidents {

		_, exists := (*unique)[item.Title]

		if !exists {
			filtered = append(filtered, item)
			(*unique)[item.Title] = true
		}
	}

	return filtered
}

func (sw *StorerWorker) buildIncidentHashset() (map[string]bool, error) {

	uniqueIncidents := make(map[string]bool)

	dbi, err := sw.storer.GetAllFromToday()
	if err != nil {
		return nil, err
	}

	for _, val := range dbi {
		uniqueIncidents[val.Title] = true
	}

	return uniqueIncidents, nil
}
