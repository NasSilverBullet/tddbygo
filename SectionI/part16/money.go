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

// Times implements Expression
func (m *Money) Times(multiplier int) Expression {
	return NewMoney(m.amount*multiplier, m.currency)
}

// Plus implements Expression
func (m *Money) Plus(added Expression) Expression {
	return NewSum(m, added)
}

// Reduce implements Expression
func (m *Money) Reduce(b *Bank, to string) *Money {
	rate := b.Rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
}
