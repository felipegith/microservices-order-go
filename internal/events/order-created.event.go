package events


type OrderCreated struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}
