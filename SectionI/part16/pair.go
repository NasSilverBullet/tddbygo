package money

type pair struct {
	from string
	to   string
}

func makePair(f, t string) pair {
	return pair{
		from: f,
		to:   t,
	}
}

func (p pair) equals(tp pair) bool {
	if p.from != tp.from {
		return false
	}
	if p.to != tp.to {
		return false
	}
	return true
}

func (p pair) hashCode() int {
	return 0
}
