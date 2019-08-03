package money

// Sum is sum
type Sum struct {
	augend Expression
	addend Expression
}

// NewSum is sum constructor
func NewSum(au, ad Expression) Expression {
	return &Sum{
		augend: au,
		addend: ad,
	}
}

// Plus implements Expression
func (sum *Sum) Plus(e Expression) Expression {
	return NewSum(sum, e)
}

// Reduce implements Expression
func (sum *Sum) Reduce(b *Bank, to string) *Money {
	a := sum.augend.Reduce(b, to).amount + sum.addend.Reduce(b, to).amount
	return NewMoney(a, to)
}

