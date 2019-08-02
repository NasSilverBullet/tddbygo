package money

// NewFranc is franc constructor
func NewFranc(a int) *Money {
	return NewMoney(a, "CHF")
}
