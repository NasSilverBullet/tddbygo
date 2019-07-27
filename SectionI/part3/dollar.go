package dollar

// Dollar is dollar
type Dollar struct {
	ammount int
}

func (d *Dollar) times(n int) *Dollar {
	return &Dollar{
		ammount: d.ammount * n,
	}
}

func (d *Dollar) equals(targetd *Dollar) bool {
	return d.ammount == targetd.ammount
}
