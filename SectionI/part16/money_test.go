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
	b := NewBank()
	if expected, actual := NewDollar(10), b.Reduce(NewDollar(5).Times(2), "USD"); !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Times(2) wont %v but got %v", expected, actual)
	}
	if expected, actual := NewDollar(15), b.Reduce(NewDollar(5).Times(3), "USD"); !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Times(3) wont %v but got %v", expected, actual)
	}
	if expected, actual := NewFranc(10), b.Reduce(NewFranc(5).Times(2), "CHF"); !expected.Equals(actual) {
		t.Errorf("NewFranc(5).Times(2) wont %v but got %v", expected, actual)
	}
	if expected, actual := NewFranc(15), b.Reduce(NewFranc(5).Times(3), "CHF"); !expected.Equals(actual) {
		t.Errorf("NewFranc(5).Times(3) wont %v but got %v", expected, actual)
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
