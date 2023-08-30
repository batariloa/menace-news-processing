package datastruct

type FetchedIncident struct {
	Id    int    `json:id`
	Title string `json:"title"`
}

type Articles struct {
	Results []FetchedIncident `json:"results"`
	Page    int               `json:"page"`
}

type FetchedResults struct {
	Articles Articles `json:"articles"`
}
