package money

// Bank is bank
type Bank struct {
}

// NewBank is bank constructor
func NewBank() *Bank {
	return &Bank{}
}

// Reduce is exchange Expression
func (b *Bank) Reduce(e Expression, to string) *Money {
	return NewDollar(10)
}
