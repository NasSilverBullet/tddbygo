package money

import (
	"testing"
)

func TestDollarMultiplication(t *testing.T) {
	five := NewDollar(5)
	if expected, actual := NewDollar(10), five.Times(2); !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Times(2) wont %v but got %v", expected, actual)
	}
	if expected, actual := NewDollar(15), five.Times(3); !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Times(3) wont %v but got %v", expected, actual)
	}

}
