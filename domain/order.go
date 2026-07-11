package domain

type Order struct {
	ID         int
	CustomerID int
	Items      []Product
	Total      float64
	Status     string
}
