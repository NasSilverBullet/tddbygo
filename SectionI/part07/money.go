package money

// Money is dollar
type Money struct {
	amount int
}

// Equals is a comparison function
// TODO 埋め込み先の型の情報が含まれているので取り除きたい
// TODO 現状の実装ではレシーバの埋め込み先の方が取得できないため、
// 本章で実装した型比較は実装できない 実装したい
func (m *Money) Equals(i interface{}) bool {
	tm := new(Money)
	switch i.(type) {
	case *Dollar:
		tm.amount = i.(*Dollar).amount
	case *Franc:
		tm.amount = i.(*Franc).amount
	default:
		return false
	}
	if m.amount != tm.amount {
		return false
	}
	return true
}
