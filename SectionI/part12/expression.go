package money

type Expression interface {
	Plus(*Money) Expression
}

func (m *Money) Plus(tm *Money) Expression {
	return NewMoney(m.amount+tm.amount, m.currency)
}
