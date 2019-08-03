package money

import "testing"

func TestMixedAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	b := NewBank()
	b.AddRate("CHF", "USD", 2)
	r := b.Reduce(fiveBucks.Plus(tenFrancs), "USD")
	if expected, actual := NewDollar(10), r; !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Times(2) wont %v but got %v", expected, actual)
	}
}
