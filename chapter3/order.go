package main

import (
	"errors"
	"time"
)

type Order struct {
	ID        int
	OrderLine []OrderLine
	TotalCost float64
	Created   time.Time
	User      string
}

func CreateOrder(id int, user string, totalCost float64) (*Order, error) {
	if id < 0 {
		return nil, errors.New("id should be a valid number")
	}
	if user == "" {
		return nil, errors.New("user is required")
	}
	if totalCost < 0 {
		return nil, errors.New("totalCost should be a valid number")
	}
	var orderLine []OrderLine
	order := Order{ID: id, OrderLine: orderLine, TotalCost: totalCost, Created: time.Now(), User: user}
	return &order, nil
}

func (o *Order) AppendOrderLine(newOrderLine OrderLine) *OrderLine {
	o.OrderLine = append(o.OrderLine, newOrderLine)
	return &newOrderLine
}

// I selected the quantity to change
func (o *Order) UpdateOrderLine(id int, quantity int) error {
	orderLines := o.OrderLine
	for _, orderLine := range orderLines {
		if orderLine.ID == id {
			orderLine.Quantity = quantity
			return nil
		}
	}
	return errors.New("The id provided is not an OrderLine")
}

func (o *Order) RemoveOrderLine(id int) *[]OrderLine {
	orderLine := o.OrderLine
	orderLine[id] = orderLine[len(orderLine)-1] //replace de element to delete with the one at the end position
	orderLine = orderLine[:len(orderLine)-1]
	return &orderLine
}

func (o *Order) TotalPriceLine() (float64, error) {
	if len(o.OrderLine) < 0 {
		return 0, errors.New("You dont have any OrderLines")
	}
	var totalPriceLine float64
	orderLines := o.OrderLine
	for _, orderLine := range orderLines {
		totalPriceLine += orderLine.UnitPrice
	}
	return totalPriceLine, nil
}
