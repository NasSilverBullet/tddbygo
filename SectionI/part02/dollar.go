package dollar

// Dollar is dollar
type Dollar struct {
	amount int
}

// Times multiplier dollar
func (d *Dollar) Times(m int) *Dollar {
	return &Dollar{
		amount: d.amount * m,
	}
}
