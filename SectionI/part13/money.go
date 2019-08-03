package money

// Money is dollar
type Money struct {
	amount   int
	currency string
}

// NewMoney is money constructor
func NewMoney(a int, c string) *Money {
	return &Money{amount: a, currency: c}
}

// Equals is a comparison function
func (m *Money) Equals(tm *Money) bool {
	if m.currency != tm.currency {
		return false
	}
	if m.amount != tm.amount {
		return false
	}
	return true
}

// Times multiplier money
func (m *Money) Times(multiplier int) *Money {
	return NewMoney(m.amount*multiplier, m.currency)
}

// Plus is money's plus
func (m *Money) Plus(added *Money) Expression {
	return NewSum(m, added)
}

// Reduce is
func (m *Money) Reduce(to string) *Money {
	return m
}
