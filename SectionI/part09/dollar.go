package money

// Dollar is dollar
type Dollar struct {
	*Money
}

// NewDollar is dollar constructor
func NewDollar(a int) *Dollar {
	return &Dollar{
		NewMoney(a, "USD"),
	}
}

// Times multiplier dollar
func (d *Dollar) Times(m int) *Dollar {
	return NewDollar(d.amount * m)
}
