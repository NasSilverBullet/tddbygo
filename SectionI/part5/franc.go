package dollar

// Franc is franc
type Franc struct {
	ammount int
}

// NewFranc is franc constructor
func NewFranc(a int) *Franc {
	return &Franc{
		ammount: a,
	}
}

func (d *Franc) times(m int) *Franc {
	return &Franc{
		ammount: d.ammount * m,
	}
}

func (d *Franc) equals(td *Franc) bool {
	return d.ammount == td.ammount
}
