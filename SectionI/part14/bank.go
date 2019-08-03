package money

// Bank is bank
type Bank struct {
}

var rates = map[pair]int{}

// NewBank is bank constructor
func NewBank() *Bank {
	return &Bank{}
}

// Reduce is exchange Expression
func (b *Bank) Reduce(e Expression, to string) Expression {
	switch e.(type) {
	case *Sum:
		sum := e.(*Sum)
		return sum.Reduce(b, to)
	case *Money:
		m := e.(*Money)
		return m.Reduce(b, to)
	default:
		return nil
	}
}

// Rate is exchange rate
func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	rate, ok := rates[makePair(from, to)]
	if !ok {
		return -1
	}
	return rate
}

// AddRate is
func (b *Bank) AddRate(from, to string, rate int) {
	rates[makePair(from, to)] = rate
}
