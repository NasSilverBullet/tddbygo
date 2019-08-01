package money

// Money is dollar
type Money struct {
	amount int
}

type money interface {
	Equals()
}

// Equals is a comparison function
func (m *Money) Equals(i interface{}) bool {
	tm := i.(Money)
	return m.amount == tm.amount
}
