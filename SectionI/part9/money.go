package money

// Money is dollar
type Money struct {
	amount   int
	currency string
}

// NewMoney is money constructor
func NewMoney(a int, c string) *Money {
	return &Money{amount: a, currency: c}
}

// Equals is a comparison function
// TODO 埋め込み先の型の情報が含まれているので取り除きたい
func (m *Money) Equals(i interface{}) bool {
	tm := new(Money)
	switch i.(type) {
	case *Dollar:
		tm.amount = i.(*Dollar).amount
		tm.currency = i.(*Dollar).currency
	case *Franc:
		tm.amount = i.(*Franc).amount
		tm.currency = i.(*Franc).currency
	default:
		return false
	}
	if m.currency != tm.currency {
		return false
	}
	if m.amount != tm.amount {
		return false
	}
	return true
}
