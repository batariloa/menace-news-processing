package alert

import (
	"time"
)

type AlertStorer interface {
	Save(*Alert) error
	GetLatestAlertDate() (time.Time, error)
}
