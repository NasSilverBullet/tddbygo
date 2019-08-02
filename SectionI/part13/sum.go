package money

// Sum is sum
type Sum struct {
	augend *Money
	addend *Money
}

// NewSum is sum constructor
func NewSum(au, ad *Money) Sum {
	return Sum{
		augend: au,
		addend: ad,
	}
}

func (sum Sum) reduce(to string) *Money {
	return NewMoney(sum.augend.amount+sum.addend.amount, to)
}
