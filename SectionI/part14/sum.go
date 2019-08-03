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

// Reduce implements Expression
func (sum *Sum) Reduce(b *Bank, to string) *Money {
	return NewMoney(sum.augend.amount+sum.addend.amount, to)
}

// Plus implements Expression
func (sum *Sum) Plus(m *Money) Expression {
	return NewSum(nil, nil)
}
