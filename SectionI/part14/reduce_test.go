package money

import "testing"

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	b := NewBank()
	reduced := b.Reduce(sum, "USD")
	if expected, actual := NewDollar(10), reduced; !expected.Equals(actual) {
		t.Errorf("NewBank().Reduce(NewDollar(5).Plus(NewDollar(5)), \"USD\") wont %v but got %v", expected, actual)
	}
}

func TestPlusReturnSum(t *testing.T) {
	five := NewDollar(5)
	r := five.Plus(five)
	sum := r.(*Sum)
	if expected, actual := five, sum.augend; !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Plus(NewDollar(5)).(Sum).augend wont %v but got %v", expected, actual)
	}
	if expected, actual := five, sum.addend; !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Plus(NewDollar(5)).(Sum).addend wont %v but got %v", expected, actual)
	}
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	b := NewBank()
	reduced := b.Reduce(sum, "USD")
	if expected, actual := NewDollar(7), reduced; !expected.Equals(actual) {
		t.Errorf("NewBank().Reduce(NewSum(NewDollar(3), NewDollar(4)), \"USD\") wont %v but got %v", expected, actual)
	}
}

func TestReduceMoney(t *testing.T) {
	b := NewBank()
	r := b.Reduce(NewDollar(1), "USD")
	if expected, actual := NewDollar(1), r; !expected.Equals(actual) {
		t.Errorf("NewBank().Reduce(NewDollar(1), \"USD\") wont %v but got %v", expected, actual)
	}
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	b := NewBank()
	b.AddRate("CHF", "USD", 2)
	r := b.Reduce(NewFranc(2), "USD")
	if expected, actual := NewDollar(1), r; !expected.Equals(actual) {
		t.Errorf("NewBank().Reduce(NewFranc(2), \"USD\") wont %v but got %v >>>> %v", expected, actual, rates)
	}
}
