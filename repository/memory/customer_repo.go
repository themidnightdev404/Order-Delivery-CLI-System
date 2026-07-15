package memory

import (
	"errors"
	"order_system/domain"
)

type MemoryCustomerRepository struct {
	customers map[int]*domain.Customer
}

func (r *MemoryCustomerRepository) GetByID(id int) (*domain.Customer, error) {
	value, exist := r.customers[id]
	if !exist {
		return nil, errors.New("customer not found")
	}
	return value, nil
}

func NewMemoryCustomerRepository() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[int]*domain.Customer),
	}
}

// Add добавляет нового клиента в репозиторий.
// Возвращает ошибку, если:
//   - передан nil;
//   - клиент с таким ID уже существует.
func (r *MemoryCustomerRepository) Add(customer *domain.Customer) error {
	// Защита от передачи nil вместо клиента.
	// Без этой проверки при обращении к customer.ID будет ошибка.
	if customer == nil {
		return errors.New("customer is nil")
	}

	// Проверяем, существует ли клиент с таким ID.
	// Второе возвращаемое значение (ok) показывает,
	// найден ключ в map или нет.
	_, ok := r.customers[customer.ID]

	// Не разрешаем добавлять клиента с уже существующим ID.
	if ok {
		return errors.New("customer already exists")
	}

	// Сохраняем указатель на клиента в map.
	// Ключом выступает уникальный ID клиента.
	r.customers[customer.ID] = customer

	// Возвращаем nil, если ошибок не возникло.
	return nil
}
