package incident

type IncidentStorer interface {
	Save(*StoredIncident) error
	GetAll() ([]*StoredIncident, error)
	GetAllFromToday() ([]*StoredIncident, error)
	SaveAll(incidents []*StoredIncident) error
}
