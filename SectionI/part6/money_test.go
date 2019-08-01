package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if expected, actual := NewDollar(10), five.Times(2); *expected != *actual {
		t.Errorf("product wont %v but got %v", expected, actual)
	}
	if expected, actual := NewDollar(15), five.Times(3); *expected != *actual {
		t.Errorf("product wont %v but got %v", expected, actual)
	}

}

func TestEquality(t *testing.T) {
	d := NewDollar(5)
	if expected, actual := true, d.Equals(NewDollar(5)); expected != actual {
		t.Errorf("Dollar{5}.equals(Dollar{5} wont %t but got %t", expected, actual)
	}
	if expected, actual := false, d.Equals(NewDollar(6)); expected != actual {
		t.Errorf("Dollar{5}.equals(Dollar{5} wont %t but got %t", expected, actual)
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	if expected, actual := NewFranc(10), five.Times(2); *expected != *actual {
		t.Errorf("product wont %v but got %v", expected, actual)
	}
	if expected, actual := NewFranc(15), five.Times(3); *expected != *actual {
		t.Errorf("product wont %v but got %v", expected, actual)
	}

}

func TestFrancEquality(t *testing.T) {
	f := NewFranc(5)
	if expected, actual := true, f.Equals(NewFranc(5)); expected != actual {
		t.Errorf("Franc{5}.equals(Franc{5} wont %t but got %t", expected, actual)
	}
	if expected, actual := false, f.Equals(NewFranc(6)); expected != actual {
		t.Errorf("Franc{5}.equals(Franc{5} wont %t but got %t", expected, actual)
	}
}
