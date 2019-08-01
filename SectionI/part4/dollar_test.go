package dollar

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if expected, actual := NewDollar(10), five.Times(2); !expected.Equals(actual) {
		t.Errorf("product wont %v but got %v", expected, actual)
	}
	if expected, actual := NewDollar(15), five.Times(3); !expected.Equals(actual) {
		t.Errorf("product wont %v but got %v", expected, actual)
	}

}

func TestEquality(t *testing.T) {
	d := NewDollar(5)
	if expected, actual := true, d.Equals(&Dollar{amount: 5}); expected != actual {
		t.Errorf("Dollar{5}.equals(Dollar{5} wont %t but got %t", expected, actual)
	}
	if expected, actual := false, d.Equals(&Dollar{amount: 6}); expected != actual {
		t.Errorf("Dollar{5}.equals(Dollar{5} wont %t but got %t", expected, actual)
	}
}
