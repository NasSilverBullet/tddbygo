package money

// Franc is franc
type Franc struct {
	*Money
}

// NewFranc is franc constructor
func NewFranc(a int) *Franc {
	return &Franc{
		NewMoney(a, "CHF"),
	}
}

// Times multiplier franc
func (f *Franc) Times(m int) *Franc {
	return NewFranc(f.amount * m)
}
