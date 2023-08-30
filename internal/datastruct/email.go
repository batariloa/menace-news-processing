package datastruct

type Email struct {
	Recipient string
	Title     string
	Location  string
}

func NewEmail(recipient string, title string, location string) *Email {

	return &Email{
		Recipient: recipient,
		Location:  location,
		Title:     title,
	}
}
