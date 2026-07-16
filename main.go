package main

import (
	"fmt"
	"order_system/domain"
	"order_system/repository/memory"
	"order_system/service"
)

func main() {
	// Тестовый клиент
	customer := domain.Customer{
		User: domain.User{
			ID:   5,
			Name: "Ars",
		},
		IsVIP:       true,
		BonusPoints: 100,
	}

	// Тестовый товар
	product := domain.Product{
		ID:    15,
		Title: "Iphone",
		Price: 100,
	}

	// Репозитории
	customerRepo := memory.NewMemoryCustomerRepository()
	orderRepo := memory.NewMemoryOrderRepository()

	// Сохраняем клиента
	err := customerRepo.Add(&customer)
	if err != nil {
		fmt.Println("Не удалось добавить клиента:", err)
		return
	}

	// Сервис заказов
	orderService := service.NewOrderService(
		customerRepo,
		orderRepo,
		service.DiscountCalculator{},
	)

	// Создаем заказ через сервис
	createdOrder, err := orderService.CreateOrder(
		1,
		customer.ID,
		[]domain.Product{product},
	)
	if err != nil {
		fmt.Println("Не удалось создать заказ:", err)
		return
	}

	// Вывод результата
	fmt.Println("Заказ создан!")
	fmt.Println("ID:", createdOrder.ID)
	fmt.Println("ID клиента:", createdOrder.CustomerID)
	fmt.Println("Товары:", createdOrder.Items)
	fmt.Println("Сумма:", createdOrder.Total)
	fmt.Println("Статус:", createdOrder.Status)
}
