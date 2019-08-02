package money

import "testing"

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	if expected, actual := NewFranc(10), five.Times(2); !expected.Equals(actual) {
		t.Errorf("NewFranc(5).Times(2) wont %v but got %v", expected, actual)
	}
	if expected, actual := NewFranc(15), five.Times(3); !expected.Equals(actual) {
		t.Errorf("NewFranc(5).Times(3) wont %v but got %v", expected, actual)
	}
}

func TestFrancEquality(t *testing.T) {
	f := NewFranc(5)
	if expected, actual := true, f.Equals(NewFranc(5)); expected != actual {
		t.Errorf("NewFranc(5).Equals(NewFranc(5)) wont %t but got %t", expected, actual)
	}
	if expected, actual := false, f.Equals(NewFranc(6)); expected != actual {
		t.Errorf("NewFranc(5).Equals(NewFranc(6)) wont %t but got %t", expected, actual)
	}
}
