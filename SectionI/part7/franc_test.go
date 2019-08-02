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
