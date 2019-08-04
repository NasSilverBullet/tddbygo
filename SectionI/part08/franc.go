package money

// Franc is franc
type Franc struct {
	*Money
}

// NewFranc is franc constructor
func NewFranc(a int) *Franc {
	return &Franc{
		&Money{
			amount: a,
		},
	}
}

// Times multiplier franc
func (f *Franc) Times(m int) *Franc {
	return &Franc{
		&Money{
			amount: f.amount * m,
		},
	}
}
