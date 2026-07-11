package service

type DiscountCalculator interface {
	Calculator(amount float64) float64
}

type RegularDiscount struct{}
type VIPDiscount struct{}
