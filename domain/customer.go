package domain

type Customer struct {
	User
	IsVIP       bool
	BonusPoints int
}
