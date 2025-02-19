package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Order struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func Create(name, email string) (*Order, error){
	order := &Order{
		Id: uuid.New().String(),
		Name: name,
		Email: email,
	}
	err := order.isValid()

	if err != nil{
		return nil, err
	}

	return order, nil
}

func (o *Order) isValid() error{
	if(len(o.Name) <= 0){
		return fmt.Errorf("Name must be greater than 0 %d", len(o.Name))
	}

	if(len(o.Email) <= 0){
		return fmt.Errorf("Email must be greater than 0 %d", len(o.Email))
	}

	return nil
}