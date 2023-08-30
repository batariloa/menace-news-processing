package worker

import (
	"fmt"
	"time"

	"github.com/batariloa/bobber/internal/datastruct"
	"github.com/batariloa/bobber/internal/service"
)

type FetcherWorker struct {
	incidentCh      *chan *[]datastruct.FetchedIncident
	IncidentService *service.IncidentService
}

func NewFetcherWorker(incidentCh *chan *[]datastruct.FetchedIncident, is *service.IncidentService) *FetcherWorker {
	return &FetcherWorker{
		incidentCh:      incidentCh,
		IncidentService: is,
	}
}

func (sw *FetcherWorker) Start() error {

	fmt.Println("Started FW")
	ticker := time.NewTicker(45 * time.Second)

	for {
		news, err := service.GetRecentArticles()

		if err != nil {

			fmt.Println("FW: ", err)
			return err
		}

		incidents := sw.IncidentService.MoldIncidents(news)

		*sw.incidentCh <- &incidents
		<-ticker.C
	}
}
