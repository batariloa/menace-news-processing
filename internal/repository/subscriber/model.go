package subscriber

type Subscriber struct {
	Id    string `json:"id" bson:"id"`
	Email string `json:"email" bson:"email"`
}

func NewSubscriber(email string) *Subscriber {
	return &Subscriber{
		Email: email,
	}
}
