package main

import (
	"fmt"

	"github.com/hgribeiro/aluno-go/order/entity"
)

func main() {
	order, err := entity.NewOrder("123", 10, 1)
	if err != nil {
		panic(err)
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Order: %+v\n", order)
}
