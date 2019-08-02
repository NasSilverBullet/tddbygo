package money

// Expression money's expression
type Expression interface {
	Plus(*Money) Expression
}

// Plus is money'plus
func (m *Money) Plus(addend *Money) Expression {
	return NewMoney(m.amount+addend.amount, m.currency)
}
