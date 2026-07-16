package service

import "order_system/domain"

type CustomerRepository interface {
	Add(customer *domain.Customer) error
	GetByID(id int) (*domain.Customer, error)
}

type OrderRepository interface {
	Add(order *domain.Order) error
	GetByID(id int) (*domain.Order, error)
}
