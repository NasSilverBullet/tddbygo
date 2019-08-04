package dollar

// Dollar is dollar
type Dollar struct {
	amount int
}

// NewDollar is dollar constructor
func NewDollar(a int) *Dollar {
	return &Dollar{
		amount: a,
	}
}

// Times multiplier dollar
func (d *Dollar) Times(m int) *Dollar {
	return &Dollar{
		amount: d.amount * m,
	}
}

// Equals is a comparison function
func (d *Dollar) Equals(td *Dollar) bool {
	return d.amount == td.amount
}
