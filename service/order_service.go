package service

import (
	"errors"
	"order_system/domain"
	repository "order_system/repository"
)

type OrderService struct {
	customerRepo repository.CustomerRepository
	orderRepo    repository.OrderRepository
	discountCalc DiscountCalculator
}

func NewOrderService(
	customerRepo repository.CustomerRepository,
	orderRepo repository.OrderRepository,
	discountCalc DiscountCalculator) *OrderService {
	return &OrderService{
		customerRepo: customerRepo,
		orderRepo:    orderRepo,
		discountCalc: discountCalc,
	}

}

func (s *OrderService) CreateOrder(
	orderID int, customerID int,
	items []domain.Product) (*domain.Order, error) {
	customer, err := s.customerRepo.GetByID(customerID)
	if err != nil {
		return nil, errors.New("customer not found")
	}
	var total float64
	for _, item := range items {
		total += item.Price
	}
	finalTotal := s.discountCalc.Calculate(customer, total)
	order := domain.Order{
		ID:         orderID,
		CustomerID: customerID,
		Items:      items,
		Total:      finalTotal,
		Status:     "New",
	}
	err = s.orderRepo.Add(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
