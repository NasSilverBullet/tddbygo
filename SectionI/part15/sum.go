package money

// Sum is sum
type Sum struct {
	augend *Money
	addend *Money
}

// NewSum is sum constructor
func NewSum(au, ad *Money) Expression {
	return &Sum{
		augend: au,
		addend: ad,
	}
}

// Plus implements Expression
func (sum *Sum) Plus(m *Money) Expression {
	return NewSum(nil, nil)
}

// Reduce implements Expression
func (sum *Sum) Reduce(b *Bank, to string) Expression {
	a := sum.augend.Reduce(b, to).(*Money).amount + sum.addend.Reduce(b, to).(*Money).amount
	return NewMoney(a, to)
}
