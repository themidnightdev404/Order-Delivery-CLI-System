package main

import (
	"fmt"
	"order_system/domain"
)

func main() {
	customer := domain.Customer{
		User: domain.User{
			ID:   1,
			Name: "Арсен",
		},
		IsVIP:       true,
		BonusPoints: 500,
	}

	laptop := domain.Product{ID: 21, Title: "MacBook", Price: 100_000}
	phone := domain.Product{ID: 22, Title: "Iphone 15", Price: 80_000}

	order := domain.Order{
		ID:         5000,
		CustomerID: customer.ID,
		Items:      []domain.Product{laptop, phone},
		Status:     "New",
	}

	for _, items := range order.Items {
		order.Total += items.Price
	}

	fmt.Printf("Заказ № %d для клиента %s на сумму %.2f руб. успешен!\n",
		order.ID, customer.Name, order.Total)
}
