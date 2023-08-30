package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/batariloa/bobber/api/server"
	"github.com/batariloa/bobber/config"
	"github.com/batariloa/bobber/internal/datastruct"
	"github.com/batariloa/bobber/internal/repository/alert"
	"github.com/batariloa/bobber/internal/repository/incident"
	"github.com/batariloa/bobber/internal/repository/subscriber"
	"github.com/batariloa/bobber/internal/service"
	"github.com/batariloa/bobber/internal/worker"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceContainer struct {
	AlertService      *service.AlertService
	SubscriberService *service.SubscriberService
	IncidentService   *service.IncidentService
	EmailService      *service.EmailService
}

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		fmt.Println("Error loading configuration: ", err)
		return
	}

	incidentCh := make(chan *[]datastruct.FetchedIncident)
	mutex := &sync.Mutex{}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Printf("Can't establish a database connection: %v", err)
		return
	}

	incidentStore := incident.NewIncidentMongoStore(client)
	alertsStore := alert.NewAlertsMongoStore(client)
	subscriberStore := subscriber.NewSubscriberMongoStore(client)

	container := ServiceContainer{}

	container.IncidentService = service.NewIncidentService()
	container.SubscriberService = service.NewSubscriberService(subscriberStore)
	container.EmailService = service.NewEmailService(config.EmailUsername, config.EmailPassword, config.EmailHost, 587)
	container.AlertService = service.NewAlertService(alertsStore, *container.SubscriberService, *container.EmailService)

	fw := worker.NewFetcherWorker(&incidentCh, container.IncidentService)
	sw := worker.NewStorerWorker(&incidentCh, incidentStore, mutex, container.AlertService)

	go func() { fw.Start() }()
	go func() { sw.Start() }()

	server := server.New(container.SubscriberService)
	server.StartHttpServer("localhost:8081")
}
