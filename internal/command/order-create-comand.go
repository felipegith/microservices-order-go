package command


type OrderCreateCommand struct {
	Name  string `json:"name"`
    Email string `json:"email"`
}