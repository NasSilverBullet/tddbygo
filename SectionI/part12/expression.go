package money

// Expression money's expression
type Expression interface {
	Plus(*Money) Expression
}

// Plus is money'plus
func (m *Money) Plus(added *Money) Expression {
	return NewMoney(m.amount+added.amount, m.currency)
}
