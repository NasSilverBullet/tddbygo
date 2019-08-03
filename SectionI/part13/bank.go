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
	switch e.(type) {
	case *Money:
		m := e.(*Money)
		return m.Reduce(to)
	case *Sum:
		sum := e.(*Sum)
		return sum.Reduce(to)
	default:
		return nil
	}
}
