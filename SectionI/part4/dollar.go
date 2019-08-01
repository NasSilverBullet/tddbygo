package dollar

// Dollar is dollar
type Dollar struct {
	ammount int
}

// NewDollar is dollar constructor
func NewDollar(a int) *Dollar {
	return &Dollar{
		ammount: a,
	}
}

func (d *Dollar) times(m int) *Dollar {
	return &Dollar{
		ammount: d.ammount * m,
	}
}

func (d *Dollar) equals(td *Dollar) bool {
	return d.ammount == td.ammount
}
