package money

// Franc is franc
type Franc struct {
	amount int
}

// NewFranc is franc constructor
func NewFranc(a int) *Franc {
	return &Franc{
		amount: a,
	}
}

// Times multiplier franc
func (f *Franc) Times(m int) *Franc {
	return &Franc{
		amount: f.amount * m,
	}
}

// Equals is a comparison function
func (f *Franc) Equals(tf *Franc) bool {
	return f.amount == tf.amount
}
