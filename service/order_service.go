package service

type OrderService struct {
	orderRepo    OrderRepository
	customerRepo CustomerRepository
}

//func (r *OrderService) CreateOrder(customerID int, products []domain.Product) (*domain.Order, error) {}
