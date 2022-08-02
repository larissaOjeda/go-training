package main

import (
	"errors"
	"time"
)

type OrderLine struct {
	ID        int
	Item      string
	Created   time.Time
	Quantity  int
	UnitPrice float64
}

func CreateOrderLine(id int, item string, quantity int, unitPrice float64) (*OrderLine, error) {
	if id < 0 {
		return nil, errors.New("id should be a valid number")
	}
	if item == "" {
		return nil, errors.New("item is required")
	}
	if quantity < 0 {
		return nil, errors.New("quantity is required")
	}
	if unitPrice < 0 {
		return nil, errors.New("unitPrice should be a valid number")
	}
	orderLine := OrderLine{ID: id, Item: item, Created: time.Now(), Quantity: quantity, UnitPrice: unitPrice}
	return &orderLine, nil
}
