package dollar

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if expected, actual := NewDollar(10), five.Times(2); !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Times(2) wont %v but got %v", expected, actual)
	}
	if expected, actual := NewDollar(15), five.Times(3); !expected.Equals(actual) {
		t.Errorf("NewDollar(5).Times(3) wont %v but got %v", expected, actual)
	}

}

func TestEquality(t *testing.T) {
	d := NewDollar(5)
	if expected, actual := true, d.Equals(NewDollar(5)); expected != actual {
		t.Errorf("NewDollar(5).equals(NewDollar(5)) wont %t but got %t", expected, actual)
	}
	if expected, actual := false, d.Equals(NewDollar(6)); expected != actual {
		t.Errorf("NewDollar(5).Equals(NewDollar(6)) wont %t but got %t", expected, actual)
	}
}
