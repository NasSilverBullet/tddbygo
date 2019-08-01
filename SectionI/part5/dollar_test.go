package dollar

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if expected, actual := NewDollar(10), five.times(2); *expected != *actual {
		t.Errorf("product wont %v but got %v", expected, actual)
	}
	if expected, actual := NewDollar(15), five.times(3); *expected != *actual {
		t.Errorf("product wont %v but got %v", expected, actual)
	}

}

func TestEquality(t *testing.T) {
	d := NewDollar(5)
	if expected, actual := true, d.equals(NewDollar(5)); expected != actual {
		t.Errorf("Dollar{5}.equals(Dollar{5} wont %t but got %t", expected, actual)
	}
	if expected, actual := false, d.equals(NewDollar(6)); expected != actual {
		t.Errorf("Dollar{5}.equals(Dollar{5} wont %t but got %t", expected, actual)
	}
}
