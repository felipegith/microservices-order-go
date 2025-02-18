package handler

import (
	"context"
	"log"
	"reflect"

	"www.github.com/felipegith/internal/command"
	"www.github.com/felipegith/internal/contract"
	"www.github.com/felipegith/internal/domain"
	"www.github.com/felipegith/internal/events"
	"www.github.com/felipegith/internal/events/helper"
)

type OrderCreateHandler struct {
	repository contract.OrderRepository
}

func Constructor(repository contract.OrderRepository) *OrderCreateHandler{
	return &OrderCreateHandler{repository: repository}
}

func (h *OrderCreateHandler) Handle(ctx context.Context, cmd command.OrderCreateCommand) (*domain.Order, error){
	order , err:= domain.Create(cmd.Name, cmd.Email)

	if err != nil{
		return nil, err
	}

	if err := h.repository.Create(ctx, order); err != nil{
		return nil, err
	}

	event := events.OrderCreated{
		 Id: order.Id,
		 Name: order.Name,
		 Email: order.Email,
	}

	err = helper.Send(reflect.TypeOf(event).Name(), event)
	
	if err != nil {
		log.Println("Error to send event", err)
		return nil, err
	}
	return order, nil
}