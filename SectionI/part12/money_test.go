package money

import "testing"

func TestEquality(t *testing.T) {
	if expected, actual := true, NewDollar(5).Equals(NewDollar(5)); expected != actual {
		t.Errorf("NewDollar(5).Equals(NewDollar(5)) wont %t but got %t", expected, actual)
	}
	if expected, actual := false, NewDollar(5).Equals(NewDollar(6)); expected != actual {
		t.Errorf("NewDollar(5).Equals(NewDollar(6)) wont %t but got %t", expected, actual)
	}
	if expected, actual := false, NewFranc(5).Equals(NewDollar(5)); expected != actual {
		t.Errorf("NewFranc(5).Equals(NewDollar(5)) wont %t but got %t", expected, actual)
	}
}

func TestMultiplication(t *testing.T) {
	if expected, actual := NewDollar(10), NewDollar(5).Times(2); !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Times(2) wont %v but got %v", expected, actual)
	}
	if expected, actual := NewDollar(15), NewDollar(5).Times(3); !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Times(3) wont %v but got %v", expected, actual)
	}
}

func TestCurrency(t *testing.T) {
	if expected, actual := "USD", NewDollar(1).currency; expected != actual {
		t.Errorf("NewDollar(1).currency wont %v but got %v", expected, actual)
	}
	if expected, actual := "CHF", NewFranc(1).currency; expected != actual {
		t.Errorf("NewFranc(1).currency wont %v but got %v", expected, actual)
	}
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	b := NewBank()
	reduced := b.Reduce(sum, "USD")
	if expected, actual := NewDollar(10), reduced; !expected.Equals(actual) {
		t.Errorf("NewDollar(5).plus(NewDollar(5)) wont %v but got %v", expected, actual)
	}
}
