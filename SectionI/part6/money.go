package money

// Money is dollar
type Money struct {
	amount int
}

// Equals is a comparison function
func (m *Money) Equals(i interface{}) bool {
	mt := new(Money)
	switch i.(type) {
	case *Dollar:
		mt.amount = i.(*Dollar).amount
	case *Franc:
		mt.amount = i.(*Franc).amount
	default:
		return false
	}
	if m.amount != mt.amount {
		return false
	}
	return true
}
