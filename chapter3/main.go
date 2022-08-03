package main

import (
	"fmt"
)

func main() {
	items := []string{"book", "pencil", "paper", "notebooks"}
	prices := []float64{0.99, 2.99, 4.89, 10.99}
	order, _ := CreateOrder(1, "andrea meza", 23.44)
	fmt.Println(order)
	for i, item := range items {
		for _, price := range prices {
			orderLine, _ := CreateOrderLine(i, item, i*2, price)
			order.AppendOrderLine(*orderLine)
			order.UpdateOrderLine(i, 6)
			order.RemoveOrderLine(0)
			order.TotalPriceLine()
		}
	}

}
