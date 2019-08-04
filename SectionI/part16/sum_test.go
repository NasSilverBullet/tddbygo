package money

import "testing"

func TestSumPlusMoney(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	b := NewBank()
	b.AddRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).Plus(fiveBucks)
	r := b.Reduce(sum, "USD")
	if expected, actual := NewDollar(15), r; !expected.Equals(actual) {
		t.Errorf("NewBank().Reduce(NewSum(fiveBucks, tenFrancs).Plus(fiveBucks)), \"USD\") wont %v but got %v", expected, actual)
	}
}

func TestSumTimes(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	b := NewBank()
	b.AddRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).Times(2)
	r := b.Reduce(sum, "USD")
	if expected, actual := NewDollar(20), r; !expected.Equals(actual) {
		t.Errorf("NewBank().Reduce(NewSum(fiveBucks, tenFrancs).Times()), \"USD\") wont %v but got %v", expected, actual)
	}
}
