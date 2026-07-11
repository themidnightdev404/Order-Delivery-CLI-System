package domain

type User struct {
	ID   int
	Name string
}

type Customer struct {
	User
	IsVIP       bool
	BonusPoints int
}
