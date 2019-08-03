package money

// NewDollar is dollar constructor
func NewDollar(a int) *Money {
	return NewMoney(a, "USD")
}
