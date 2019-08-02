package money

// Expression money's expression
type Expression interface {
	Plus(*Money) Expression
	Reduce(*Bank, string) *Money
}

// Plus implements Expression
func (Sum) Plus(*Money) Expression {
	return nil
}

// Plus is money's plus
func (m *Money) Plus(added *Money) Expression {
	return NewSum(m, added)
}

// Reduce is
func (m *Money) Reduce(b *Bank, to string) *Money {
	rate := b.Rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
}

// Reduce is
func (sum Sum) Reduce(b *Bank, to string) *Money {
	return NewMoney(sum.augend.amount+sum.addend.amount, to)
}
