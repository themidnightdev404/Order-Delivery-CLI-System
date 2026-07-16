package memory

import (
	"errors"
	"order_system/domain"
)

type MemoryOrderRepository struct {
	orders map[int]*domain.Order
}

func (r *MemoryOrderRepository) GetByID(id int) (*domain.Order, error) {
	value, exist := r.orders[id]
	if !exist {
		return nil, errors.New("order not found")
	}
	return value, nil
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		orders: make(map[int]*domain.Order),
	}
}

func (r *MemoryOrderRepository) Add(order *domain.Order) error {
	if order == nil {
		return errors.New("order is nil")
	}

	_, ok := r.orders[order.ID]
	if ok {
		return errors.New("order already exists")
	}
	r.orders[order.ID] = order
	return nil
}
