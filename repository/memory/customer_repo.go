package memory

import (
	"errors"
	"order_system/domain"
)

type MemoryCustomerRepository struct {
	customers map[int]*domain.Customer
}

func (r *MemoryCustomerRepository) GetByID(id int) (*domain.Customer, error) {
	value, exists := r.customers[id]
	if !exists {
		return nil, errors.New("not found")
	}
	return value, nil
}

type MemoryOrderRepository struct {
	orders map[int]*domain.Order
}

//func (r *MemoryOrderRepository) Save(order *domain.Order) error {}

//func (r *MemoryOrderRepository) GetByID(id int) (*domain.Order, error) {}
