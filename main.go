package main

import (
	"fmt"
	"order_system/domain"
	"order_system/repository/memory"
)

func main() {

	customer := domain.Customer{
		User: domain.User{
			ID:   5,
			Name: "Ars",
		},
		IsVIP:       true,
		BonusPoints: 100,
	}

	product := domain.Product{
		ID:    15,
		Title: "Iphone",
		Price: 100,
	}

	order := domain.Order{
		ID:         4,
		CustomerID: 5,
		Items: []domain.Product{
			product,
		},
		Total:  100,
		Status: "Paid",
	}

	// Создание репозиториев

	customerRepo := memory.NewMemoryCustomerRepository()
	orderRepo := memory.NewMemoryOrderRepository()

	// Проверка CustomerRepository

	err := customerRepo.Add(&customer)
	if err != nil {
		fmt.Println("Не удалось добавить клиента:", err)
		return
	}

	// Проверяем защиту от дублирования ID.
	err = customerRepo.Add(&customer)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	foundCustomer, err := customerRepo.GetByID(5)
	if err != nil {
		fmt.Println("Не удалось найти клиента:", err)
		return
	}

	fmt.Println("Клиент найден!")
	fmt.Println("ID:", foundCustomer.ID)
	fmt.Println("Имя:", foundCustomer.Name)
	fmt.Println("ВИП:", foundCustomer.IsVIP)
	fmt.Println("Бонусы:", foundCustomer.BonusPoints)

	_, err = customerRepo.GetByID(999)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	// Проверка OrderRepository

	err = orderRepo.Add(&order)
	if err != nil {
		fmt.Println("Не удалось добавить заказ:", err)
		return
	}

	// Проверяем защиту от дублирования ID.
	err = orderRepo.Add(&order)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	foundOrder, err := orderRepo.GetByID(4)
	if err != nil {
		fmt.Println("Не удалось найти заказ:", err)
		return
	}

	fmt.Println("Заказ найден!")
	fmt.Println("ID:", foundOrder.ID)
	fmt.Println("ID клиента:", foundOrder.CustomerID)
	fmt.Println("Товары:", foundOrder.Items)
	fmt.Println("Сумма:", foundOrder.Total)
	fmt.Println("Статус:", foundOrder.Status)

	_, err = orderRepo.GetByID(999)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

}
