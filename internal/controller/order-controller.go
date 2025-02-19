package controller

import (
	"encoding/json"
	"net/http"

	"www.github.com/felipegith/internal/command"
	"www.github.com/felipegith/internal/handler"
)


type OrderController struct {
	createHandler *handler.OrderCreateHandler
}

func NewOrderController(createHandler *handler.OrderCreateHandler) *OrderController {
	return &OrderController{
		createHandler: createHandler,
	}
}

func (c *OrderController) Create(w http.ResponseWriter, r *http.Request) {
	var command command.OrderCreateCommand
	
	if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := c.createHandler.Handle(r.Context(), command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}