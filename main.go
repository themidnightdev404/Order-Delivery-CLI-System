package main

import (
	"fmt"
	"order_system/domain"
	"order_system/repository/memory"
)

func main() {
	repo := memory.NewMemoryCustomerRepository()
	customer := domain.Customer{
		User: domain.User{
			ID:   5,
			Name: "Arsen",
		},
		IsVIP:       true,
		BonusPoints: 100,
	}
	err := repo.Add(&customer)
	if err != nil {
		fmt.Println("Не удалось добавить клиента:", err)
		return
	}
}
