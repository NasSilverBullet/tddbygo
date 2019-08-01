package money

// Dollar is dollar
type Dollar struct {
	amount int
}

// NewDollar is dollar constructor
func NewDollar(a int) *Dollar {
	return &Dollar{
		amount: a,
	}
}

// Times multiplier dollar
func (d *Dollar) Times(m int) *Dollar {
	return &Dollar{
		amount: d.amount * m,
	}
}

// Equals is a comparison function
func (d *Dollar) Equals(td *Dollar) bool {
	return d.amount == td.amount
}

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
