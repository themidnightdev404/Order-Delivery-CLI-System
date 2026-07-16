package service

import "order_system/domain"

type DiscountCalculator struct{}

const VIPDiscountPercent = 10

func (r *DiscountCalculator) Calculate(customer *domain.Customer, sum float64) float64 {
	if customer.IsVIP {
		discount := sum * VIPDiscountPercent / 100
		return sum - discount
	}
	return sum
}
