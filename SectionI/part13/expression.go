package money

// Expression money's expression
type Expression interface {
	Plus(*Money) Expression
}

// Plus implements Expression
func (Sum) Plus(*Money) Expression {
	return nil
}

// Plus is money's plus
func (m *Money) Plus(added *Money) Expression {
	return NewSum(m, added)
}
