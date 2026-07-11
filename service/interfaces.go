package service

import "order_system/domain"

type CustomerRepository interface {
	GetByID(id int) (*domain.Customer, error)
}

type OrderRepository interface {
	Save(order *domain.Order) error
	GetByID(id int) (*domain.Order, error)
	GetAllByCustomerID(customerID int) ([]domain.Order, error)
}
