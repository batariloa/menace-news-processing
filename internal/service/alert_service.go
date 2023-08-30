package service

import (
	"fmt"

	"github.com/batariloa/bobber/internal/datastruct"
	"github.com/batariloa/bobber/internal/repository/alert"
	"github.com/batariloa/bobber/internal/repository/incident"
)

type AlertService struct {
	store             alert.AlertStorer
	SubscriberService SubscriberService
	EmailService      EmailService
}

func NewAlertService(store alert.AlertStorer, s SubscriberService, e EmailService) *AlertService {

	return &AlertService{
		store:             store,
		SubscriberService: s,
		EmailService:      e,
	}
}

func (a *AlertService) AlertUsers(i *incident.StoredIncident) {

	subscribers, err := a.SubscriberService.GetAllSubscribers()

	if err != nil {
		fmt.Println("There was an error fetching subscribers")
		return
	}

	genLoc := "TODO Gen loc"
	for _, s := range subscribers {

		email := datastruct.NewEmail(s.Email, i.Title, genLoc)
		err := a.EmailService.SendEmail(email)

		if err != nil {
			fmt.Println("Error sending email.", err)
			return
		}

		fmt.Println("Sent email: ", email)
	}
}
