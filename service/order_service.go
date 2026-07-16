package service

import repository "order_system/repository"

type OrderService struct {
	customerRepo repository.CustomerRepository
	orderRepo    repository.OrderRepository
	discountCalc DiscountCalculator
}

func NewOrderService(customerRepo repository.CustomerRepository, orderRepo repository.OrderRepository, discountCalc DiscountCalculator) *OrderService {
	return &OrderService{
		customerRepo: customerRepo,
		orderRepo:    orderRepo,
		discountCalc: discountCalc,
	}

}
